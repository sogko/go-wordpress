# List of WP-API endpoints

List of WP-API REST endpoints and implementation status

## Attachments / Media

- [x] `GET /media`
- [x] `POST /media`
- [x] `GET /media/[id]`
- [ ] `PUT /media/[id]` (Unknown if supported)
- [x] `DELETE /media/[id]`  (requires `define( 'MEDIA_TRASH', true );` in `wp_config.php`, see: https://github.com/WP-API/WP-API/issues/1493)

## Comments

- [x] `GET    /comments`
- [ ] `POST   /comments` (Implemented but untested)
- [x] `GET    /comments/[id]`
- [ ] `PUT    /comments/[id]`  (Implemented but untested)
- [ ] `DELETE /comments/[id]`  (Implemented but untested)

## Meta

- [x] `GET    /[parent_base]/[parent_id]/meta`
- [x] `POST   /[parent_base]/[parent_id]/meta`
- [x] `GET    /[parent_base]/[parent_id]/meta/[id]`
- [x] `PUT    /[parent_base]/[parent_id]/meta/[id]`
- [x] `DELETE /[parent_base]/[parent_id]/meta/[id]`

`[parent_base] = "posts" | "pages"`

### Meta Posts

- [x] `GET    /posts/[post_id]/meta`
- [x] `POST   /posts/[post_id]/meta`
- [x] `GET    /posts/[post_id]/meta/[id]`
- [x] `PUT    /posts/[post_id]/meta/[id]`
- [x] `DELETE /posts/[post_id]/meta/[id]`

### Meta Pages

- [x] `GET    /pages/[post_id]/meta`
- [x] `POST   /pages/[post_id]/meta`
- [x] `GET    /pages/[post_id]/meta/[id]`
- [x] `PUT    /pages/[post_id]/meta/[id]`
- [x] `DELETE /pages/[post_id]/meta/[id]`

## Post Statuses

- [x] `GET    /statuses`
- [x] `GET    /statuses/[slug]`

## Post Types

- [x] `GET    /types`
- [x] `GET    /types/[slug]`

## Posts

- [x] `GET    /posts`
- [x] `POST   /posts`
- [x] `GET    /posts/[id]`
- [x] `PUT    /posts/[id]`
- [x] `DELETE /posts/[id]`

## Pages

- [x] `GET    /pages`
- [x] `POST   /pages`
- [x] `GET    /pages/[id]`
- [x] `PUT    /pages/[id]`
- [x] `DELETE /pages/[id]`

## Post Terms

- [x] `GET    /[post_base]/[post_id]/terms/[tax_base]`
- [x] `GET    /[post_base]/[post_id]/terms/[tax_base]/[term_id]`
- [x] `POST   /[post_base]/[post_id]/terms/[tax_base]/[term_id]`
- [x] `DELETE /[post_base]/[post_id]/terms/[tax_base]/[term_id]`

`[post_base] = "posts"`
`[tax_base] = "tag" | "category"`

### Post Terms Tag

- [x] `GET    /posts/[post_id]/terms/tag`
- [x] `GET    /posts/[post_id]/terms/tag/[term_id]`
- [x] `POST   /posts/[post_id]/terms/tag/[term_id]`
- [x] `DELETE /posts/[post_id]/terms/tag/[term_id]`

### Post Terms Category

- [x] `GET    /posts/[post_id]/terms/category`
- [x] `GET    /posts/[post_id]/terms/category/[term_id]`
- [x] `POST   /posts/[post_id]/terms/category/[term_id]`
- [x] `DELETE /posts/[post_id]/terms/category/[term_id]`

## Revisions

- [x] `GET    /[parent_base]/[parent_id]/revisions`
- [x] `GET    /[parent_base]/[parent_id]/revisions/[id]`
- [x] `DELETE /[parent_base]/[parent_id]/revisions/[id]`

`[parent_base] = "posts" | "pages"`

### Revisions Posts

- [x] `GET    /posts/[parent_id]/revisions`
- [x] `GET    /posts/[parent_id]/revisions/[id]`
- [x] `DELETE /posts/[parent_id]/revisions/[id]`

### Revisions Pages

- [x] `GET    /pages/[parent_id]/revisions`
- [x] `GET    /pages/[parent_id]/revisions/[id]`
- [x] `DELETE /pages/[parent_id]/revisions/[id]`

## Taxonomies

- [x] `GET    /taxonomies`
- [x] `GET    /taxonomies/[slug]`

## Terms

- [x] `GET    /terms/[tax_base]`
- [x] `POST   /terms/[tax_base]`
- [x] `GET    /terms/[tax_base]/[id]`
- [x] `PUT    /terms/[tax_base]/[id]`
- [x] `DELETE /terms/[tax_base]/[id]`

`[tax_base] = "tag" | "category"`

### Tag Terms

- [x] `GET    /terms/tag`
- [x] `POST   /terms/tag`
- [x] `GET    /terms/tag/[id]`
- [x] `PUT    /terms/tag/[id]`
- [x] `DELETE /terms/tag/[id]`

### Category Terms

- [x] `GET    /terms/category`
- [x] `POST   /terms/category`
- [x] `GET    /terms/category/[id]`
- [x] `PUT    /terms/category/[id]`
- [x] `DELETE /terms/category/[id]`

## Users

- [x] `GET    /users`
- [x] `POST   /users`
- [x] `GET    /users/[id]`
- [x] `PUT    /users/[id]`
- [x] `DELETE /users/[id]`
- [x] `GET    /users/me`


