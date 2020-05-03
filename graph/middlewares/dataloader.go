package middlewares

import (
    "context"
    "github.com/go-pg/pg/v9"
    "net/http"
    "time"

    "github.com/vickywane/event-server/graph/dataloaders"
    "github.com/vickywane/event-server/graph/model"
)

const loaderKey = "userLoader"

func DataLoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userLoader := dataloaders.UserLoader{
            MaxBatch: 100,
            Wait:     1 * time.Millisecond,
            Fetch: func(ids []int) ([]*model.User, []error) {
                var users []*model.User

                err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
                if err != nil {
                    return nil, []error{err}
                }
                return users, nil
            },
        }

        ctx := context.WithValue(r.Context(), loaderKey, &userLoader)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func getUserLoader(ctx context.Context) *dataloaders.UserLoader {
    return ctx.Value(loaderKey).(*dataloaders.UserLoader)
}
