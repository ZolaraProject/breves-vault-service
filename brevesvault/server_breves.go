package brevesvault

import (
	"context"

	"database/sql"
	"errors"
	"fmt"
	logger "github.com/ZolaraProject/library/logger"
	"math/rand"
	"strings"
	"time"

	. "github.com/ZolaraProject/breves-vault-service/brevesvaultrpc"
)

var (
	db *sql.DB
)

const (
	VideoUrlInject         = "/video/%s"
	NotUnderstoodThreshold = 5
)

func (*server) GetUserVideos(ctx context.Context, req *UserVideoRequest) (*UserVideoList, error) {
	db, err := sql.Open("postgres", DbUrl())
	if err != nil {
		logger.Err("Open error : %v", err)
		return nil, err
	}

	if req.GetUserId() == 0 {
		logger.Err("user id is mandatory")
		return nil, errors.New("user id is mandatory")
	}

	query := `
	SELECT v.id, v.title, v.subtitle, v.likes, l.name AS language, a.name AS action,
	CASE 
		WHEN ARRAY_AGG(ui.interest_id) && ARRAY_AGG(vc.interest_id) THEN true 
		ELSE false 
	END AS interest_match
	FROM users u
	INNER JOIN user_profiles up ON up.user_id = u.id
	INNER JOIN user_activity ua ON ua.user_id = u.id
	INNER JOIN user_interests ui ON ui.user_id = u.id
	INNER JOIN video_category vc ON vc.action_id = up.action_id
	INNER JOIN videos v ON v.language_id = ua.current_language_id AND v.id = vc.video_id
	INNER JOIN languages l ON l.id = v.language_id
	INNER JOIN actions a ON a.id = vc.action_id
	WHERE u.id = $1
	GROUP BY v.id, l.name, a.name
	`

	totalQuery := `SELECT COUNT(*)
	FROM (
		SELECT v.id
		FROM users u
		INNER JOIN user_profiles up ON up.user_id = u.id
		INNER JOIN user_activity ua ON ua.user_id = u.id
		INNER JOIN user_interests ui ON ui.user_id = u.id
		INNER JOIN video_category vc ON vc.action_id = up.action_id
		INNER JOIN videos v ON v.language_id = ua.current_language_id AND v.id = vc.video_id
		INNER JOIN languages l ON l.id = v.language_id
		INNER JOIN actions a ON a.id = vc.action_id
		WHERE u.id = $1
		GROUP BY v.id, l.name, a.name
	)
	`

	rows, err := db.Query(query, req.GetUserId())
	if err != nil {
		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}
	defer rows.Close()

	res := []*UserVideoInList{}
	interestingVideos := []*UserVideoInList{}
	otherVideos := []*UserVideoInList{}

	for rows.Next() {
		var (
			id            string
			title         string
			subtitle      string
			likes         int64
			language      string
			action        string
			interestMatch bool
		)

		err := rows.Scan(&id, &title, &subtitle, &likes, &language, &action, &interestMatch)
		if err != nil {
			logger.Err("failed to scan row: %s", err)
			return nil, fmt.Errorf("failed to scan row: %s", err)
		}

		userVideoInList := &UserVideoInList{
			Title:    title,
			Subtitle: subtitle,
			Likes:    likes,
			Language: Language(Language_value[strings.ToUpper(language)]),
			Action:   action,
			VideoUrl: fmt.Sprintf(VideoUrlInject, id),
			VideoId:  id,
		}

		if interestMatch {
			interestingVideos = append(interestingVideos, userVideoInList)
		} else {
			otherVideos = append(otherVideos, userVideoInList)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(interestingVideos); i++ {
		j := r.Intn(i + 1)
		interestingVideos[i], interestingVideos[j] = interestingVideos[j], interestingVideos[i]
	}
	for i := 0; i < len(otherVideos); i++ {
		j := r.Intn(i + 1)
		otherVideos[i], otherVideos[j] = otherVideos[j], otherVideos[i]
	}

	res = append(res, interestingVideos...)
	res = append(res, otherVideos...)

	var total int64 = 0
	if err := db.QueryRow(totalQuery, req.GetUserId()).Scan(&total); err != nil {
		logger.Err("failed to execute total query: %s", err)
		return nil, fmt.Errorf("failed to execute total query: %s", err)
	}

	return &UserVideoList{
		UserVideos: res,
		Total:      total,
	}, nil
}

func (s *server) UpdateUserVideo(ctx context.Context, req *UpdateUserVideoRequest) (*UserVideoList, error) {
	db, err := sql.Open("postgres", DbUrl())
	if err != nil {
		logger.Err("Open error : %v", err)
		return nil, err
	}

	isRelated, err := checkUserVideoRelation(db, req.GetUserId(), req.GetVideoId())
	if err != nil {
		logger.Err("failed to check user video relation: %s", err)
		return nil, fmt.Errorf("failed to check user video relation: %s", err)
	}
	if !isRelated {
		logger.Warn("user %d is not related to video %d", req.GetUserId(), req.GetVideoId())
		return nil, errors.New("user is not related to video")
	}

	userVideoQuery := `INSERT INTO user_videos (user_id, video_id, not_understood) VALUES ($1, $2, TRUE) ON CONFLICT (user_id, video_id) DO UPDATE SET not_understood = TRUE
	`

	_, err = db.Exec(userVideoQuery, req.GetUserId(), req.GetVideoId())
	if err != nil {
		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	adjusted, err := checkUserUnderstanding(db, req.GetVideoId(), req.GetUserId())
	if err != nil {
		logger.Err("failed to check user understanding: %s", err)
		return nil, fmt.Errorf("failed to check user understanding: %s", err)
	}

	if adjusted {
		logger.Debug("User level adjusted. Getting new video list")
		newVideoList, err := s.GetUserVideos(ctx, &UserVideoRequest{UserId: req.GetUserId()})
		if err != nil {
			logger.Err("failed to get user videos: %s", err)
			return nil, fmt.Errorf("failed to get user videos: %s", err)
		}
		return &UserVideoList{
			UserVideos: newVideoList.GetUserVideos(),
			Total:      newVideoList.GetTotal(),
		}, nil
	}

	return nil, nil
}

func checkUserUnderstanding(db *sql.DB, videoId int64, userId int64) (bool, error) {
	notUnderstoodCountQuery := `
	WITH video_details AS (
        SELECT v.language_id, vc.level_id, vc.action_id
        FROM videos v
        LEFT JOIN video_category vc ON vc.video_id = v.id
        WHERE v.id = $1
    )
	SELECT COUNT(*)
	FROM user_videos uv
    INNER JOIN videos v ON v.id = uv.video_id
    LEFT JOIN video_category vc ON vc.video_id = v.id
    INNER JOIN video_details vd ON vd.language_id = v.language_id AND vd.level_id = vc.level_id AND vd.action_id = vc.action_id
    WHERE uv.user_id = $2 AND uv.not_understood = TRUE AND v.language_id = vd.language_id AND vc.level_id = vd.level_id`

	var notUnderstoodCount int64
	if err := db.QueryRow(notUnderstoodCountQuery, videoId, userId).Scan(&notUnderstoodCount); err != nil {
		logger.Err("failed to execute query: %s", err)
		return false, fmt.Errorf("failed to execute query: %s", err)
	}

	if notUnderstoodCount >= NotUnderstoodThreshold {
		logger.Debug("The user reached the threshold of not understanding the video. Updating the user profile")
		adjusted, err := adjustUserLevel(db, videoId, userId)
		if err != nil {
			logger.Err("failed to adjust user level: %s", err)
			return false, fmt.Errorf("failed to adjust user level: %s", err)
		}
		return adjusted, nil
	}
	return false, nil
}

func adjustUserLevel(db *sql.DB, videoId int64, userId int64) (bool, error) {
	getLevelQuery := `
	WITH video_details AS (
		SELECT v.language_id, vc.level_id, vc.action_id
		FROM videos v
		LEFT JOIN video_category vc ON vc.video_id = v.id
		WHERE v.id = $1
	)
	SELECT lv.id, up.language_id, up.action_id
	FROM user_profiles up
	INNER JOIN video_details vd ON vd.language_id = up.language_id AND vd.level_id = up.level_id AND vd.action_id = up.action_id
	INNER JOIN levels lv ON lv.id = up.level_id
	WHERE up.user_id = $2
	`

	var (
		levelId    int64
		languageId int64
		actionId   int64
	)
	if err := db.QueryRow(getLevelQuery, videoId, userId).Scan(&levelId, &languageId, &actionId); err != nil {
		logger.Err("failed to execute query: %s", err)
		return false, fmt.Errorf("failed to execute query: %s", err)
	}

	if levelId > 1 {
		updateLevelQuery := `UPDATE user_profiles SET level_id = level_id - 1 WHERE user_id = $1 AND language_id = $2 AND action_id = $3`
		if _, err := db.Exec(updateLevelQuery, userId, languageId, actionId); err != nil {
			logger.Err("failed to execute query: %s", err)
			return false, fmt.Errorf("failed to execute query: %s", err)
		}
		return true, nil
	}

	logger.Debug("The user is already at the lowest level")
	return false, nil
}

func (s *server) LikeVideo(ctx context.Context, req *LikeVideoRequest) (*CreateResponse, error) {
	db, err := sql.Open("postgres", DbUrl())
	if err != nil {
		logger.Err("Open error : %v", err)
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		logger.Err("failed to begin transaction: %s", err)
		return nil, fmt.Errorf("failed to begin transaction: %s", err)
	}
	defer tx.Rollback()

	stmt1, err := tx.Prepare("UPDATE videos SET likes = likes + 1 WHERE id = $1;")
	if err != nil {
		logger.Err("failed to prepare query: %s", err)
		return nil, fmt.Errorf("failed to prepare query: %s", err)
	}
	defer stmt1.Close()

	stmt2, err := tx.Prepare("INSERT INTO user_videos (user_id, video_id, liked) VALUES ($1, $2, TRUE) ON CONFLICT (user_id, video_id) DO UPDATE SET liked = TRUE")
	if err != nil {
		logger.Err("failed to prepare query: %s", err)
		return nil, fmt.Errorf("failed to prepare query: %s", err)
	}
	defer stmt2.Close()

	for _, video := range req.GetVideoLikes() {
		isRelated, err := checkUserVideoRelation(db, req.GetUserId(), video.GetId())
		if err != nil {
			logger.Err("failed to check user video relation: %s", err)
			return nil, fmt.Errorf("failed to check user video relation: %s", err)
		}
		if !isRelated {
			logger.Warn("user %d is not related to video %d", req.GetUserId(), video.GetId())
			continue
		}

		_, err = stmt1.Exec(video.GetId())
		if err != nil {
			logger.Err("failed to execute stmt1 query: %s", err)
			return nil, fmt.Errorf("failed to execute stmt1 query: %s", err)
		}

		_, err = stmt2.Exec(req.GetUserId(), video.GetId())
		if err != nil {
			logger.Err("failed to execute stmt2 query: %s", err)
			return nil, fmt.Errorf("failed to execute stmt2 query: %s", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		logger.Err("failed to commit transaction: %s", err)
		return nil, fmt.Errorf("failed to commit transaction: %s", err)
	}

	return &CreateResponse{
		Message: "Videos liked successfully",
	}, nil
}

func checkUserVideoRelation(db *sql.DB, userId int64, videoId int64) (bool, error) {
	query := `SELECT COUNT(*)
	FROM videos v
	LEFT JOIN video_category vc ON vc.video_id = v.id
	INNER JOIN user_profiles up ON up.language_id = v.language_id AND up.actions_levels_id = vc.actions_levels_id
	INNER JOIN languages l ON l.id = v.language_id
	JOIN actions_levels al ON vc.actions_levels_id = al.id
	INNER JOIN levels lv ON lv.id = al.level_id
	INNER JOIN actions a ON a.id = al.action_id
	WHERE up.user_id = $1 AND v.id = $2
	`

	var count int64
	if err := db.QueryRow(query, userId, videoId).Scan(&count); err != nil {
		logger.Err("failed to execute query: %s", err)
		return false, fmt.Errorf("failed to execute query: %s", err)
	}

	return count > 0, nil
}

func (*server) CreateVideo(ctx context.Context, req *CreateVideoRequest) (*CreateResponse, error) {
	db, err := sql.Open("postgres", DbUrl())
	if err != nil {
		logger.Err("Open error : %v", err)
		return nil, err
	}

	videoQuery := `INSERT INTO videos (id, title, subtitle, likes, language, level, action_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err = db.QueryRow(videoQuery, req.Id, req.Title, req.Subtitle, req.Likes, req.Language, req.Level, req.ActionId).Scan()
	if err != nil {
		logger.Err("failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	return &CreateResponse{
		CreatedId: req.Id,
		Message:   "Video created successfully",
	}, nil
}
