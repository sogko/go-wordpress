# List of WP-API endpoints

List of WP-API REST endpoints and implementation status

## Attachments / Media

- [ ] `POST /media`
- [ ] `GET /media`
- [ ] `GET /media/[id]`
- [ ] `PUT /media/[id]` ?
- [ ] `DELETE /media/[id]`  (requires `define( 'MEDIA_TRASH', true );` in `wp_config.php`, see: https://github.com/WP-API/WP-API/issues/1493)

## Comments

- [x] `GET    /comments`
- [x] `POST   /comments` (Untested)
- [x] `GET    /comments/[id]`
- [x] `PUT    /comments/[id]`  (Untested)
- [x] `DELETE /comments/[id]`  (Untested)

## Meta

- [ ] `GET    /[parent_base]/[parent_id]/meta`
- [ ] `POST   /[parent_base]/[parent_id]/meta`
- [ ] `GET    /[parent_base]/[parent_id]/meta/[id]`
- [ ] `PUT    /[parent_base]/[parent_id]/meta/[id]`
- [ ] `DELETE /[parent_base]/[parent_id]/meta/[id]`

`[parent_base] = "posts" | "pages"`

### Meta Posts

- [x] `GET    /posts/[post_id]/meta`
- [x] `POST   /posts/[post_id]/meta`
- [x] `GET    /posts/[post_id]/meta/[id]`
- [x] `PUT    /posts/[post_id]/meta/[id]`
- [x] `DELETE /posts/[post_id]/meta/[id]`

### Meta Pages

- [ ] `GET    /pages/[post_id]/meta`
- [ ] `POST   /pages/[post_id]/meta`
- [ ] `GET    /pages/[post_id]/meta/[id]`
- [ ] `PUT    /pages/[post_id]/meta/[id]`
- [ ] `DELETE /pages/[post_id]/meta/[id]`

## Post Statuses

- [ ] `GET    /statuses`
- [ ] `GET    /statuses/[slug]`

## Post Types

- [ ] `GET    /types`
- [ ] `GET    /types/[slug]`

## Posts

- [x] `GET    /posts`
- [x] `POST   /posts`
- [x] `GET    /posts/[id]`
- [x] `PUT    /posts/[id]`
- [x] `DELETE /posts/[id]`

## Post Terms

- [ ] `GET    /[post_base]/[post_id]/terms/[tax_base]`
- [ ] `GET    /[post_base]/[post_id]/terms/[tax_base]/[term_id]`
- [ ] `POST   /[post_base]/[post_id]/terms/[tax_base]/[term_id]`
- [ ] `DELETE /[post_base]/[post_id]/terms/[tax_base]/[term_id]`

`[post_base] = "posts"`
`[tax_base] = "tag" | "category"`

## Revisions

- [ ] `GET    /[parent_base]/[parent_id]/revisions`
- [ ] `GET    /[parent_base]/[parent_id]/revisions/[id]`
- [ ] `DELETE /[parent_base]/[parent_id]/revisions/[id]`

`[parent_base] = "posts" | "pages"`

### Revisions Posts

- [x] `GET    /posts/[parent_id]/revisions`
- [x] `GET    /posts/[parent_id]/revisions/[id]`
- [x] `DELETE /posts/[parent_id]/revisions/[id]`

### Revisions Pages

- [ ] `GET    /pages/[parent_id]/revisions`
- [ ] `GET    /pages/[parent_id]/revisions/[id]`
- [ ] `DELETE /pages/[parent_id]/revisions/[id]`

## Taxonomies

- [x] `GET    /taxonomies`
- [x] `GET    /taxonomies/[slug]`

## Terms

- [ ] `GET    /terms/[tax_base]`
- [ ] `POST   /terms/[tax_base]`
- [ ] `GET    /terms/[tax_base]/[id]`
- [ ] `PUT    /terms/[tax_base]/[id]`
- [ ] `DELETE /terms/[tax_base]/[id]`

`[tax_base] = "tag" | "category"`

### Tag Terms

- [ ] `GET    /terms/tag`
- [ ] `POST   /terms/tag]`
- [ ] `GET    /terms/tag]/[id]`
- [ ] `PUT    /terms/tag]/[id]`
- [ ] `DELETE /terms/tag/[id]`

### Category Terms

- [ ] `GET    /terms/category`
- [ ] `POST   /terms/category]`
- [ ] `GET    /terms/category]/[id]`
- [ ] `PUT    /terms/category]/[id]`
- [ ] `DELETE /terms/category/[id]`

## Users

- [ ] `GET    /users`
- [ ] `POST   /users`
- [ ] `GET    /users/[id]`
- [ ] `PUT    /users/[id]`
- [ ] `DELETE /users/[id]`
- [ ] `GET    /users/me`


