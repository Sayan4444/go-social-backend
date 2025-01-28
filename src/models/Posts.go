package models

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// type PostStore struct {
// 	db *sql.DB
// }

// func (ps *PostStore) Create(ctx context.Context, post *Post) error {
// 	query := `
// 		INSERT INTO posts (content, title, user_id, tags)
// 		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

// 	err := ps.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (ps *PostStore) GetByID(ctx context.Context, postID int64) (*Post, error) {
// 	query := `SELECT id, content, title, user_id, tags, created_at, updated_at FROM posts WHERE id = $1`
// 	var post Post
// 	err := ps.db.QueryRowContext(ctx, query, postID).Scan(
// 		&post.ID,
// 		&post.Content,
// 		&post.Title,
// 		&post.UserID,
// 		pq.Array(&post.Tags),
// 		&post.CreatedAt,
// 		&post.UpdatedAt,
// 	)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, ErrNotFound
// 		}
// 		return nil, err
// 	}
// 	return &post, nil
// }
