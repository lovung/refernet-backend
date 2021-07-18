# refernet-backend
Backend for ReferNet

## DBML

```
// Phase 1: MVP
Table users {
  id varchar [pk]
  name varchar
  email varchar
  phone varchar
  intro varchar
  github_profile varchar
  status varchar
}


Table work_experiences {
  id varchar  [pk]
  user_id varchar [ref: > users.id]
  title varchar
  type employment_type
  company_id varchar [ref: > companies.id]
  location varchar
  start_date date
  end_date date
  description varchar
}

Table companies {
  id varchar [pk]
  name varchar
  logo_url varchar
  overview varchar
  website varchar
  industry varchar
  size company_size
  founded_at varchar
}

Table jobs {
  id varchar [pk]
  owner_id varchar [ref: > users.id]
  title varchar
  location varchar
  description varchar
  type employment_types
  company_id varchar [ref: > companies.id]
}

Table jobs_related_skills {
  id varchar [pk]
  job_id varchar [ref: > jobs.id]
  skill_id varchar [ref: > skills.id]
}

Table company_locations {
  id varchar [pk]
  company_id varchar [ref: > companies.id] 
  name varchar
  address varchar
}

Table skills {
  id varchar [pk]
  name varchar
  dark_logo_url varchar
  light_logo_url varchar
  parent_id varchar [ref: > skills.id]
}


Table work_experiences_skills {
  id varchar [pk]
  skill_id varchar [ref: > skills.id]
  work_experience_id varchar [ref: > work_experiences.id]
  percent int
}

// Phase 2: Chat feature
// Table chat_channels {
//   id varchar [pk]
//   host_id varchar [ref: > users.id]
//   participant_id varchar [ref: > users.id]
// }

// Table chat_messages {
//   id varchar [pk]
//   from_user_id varchar [ref: > users.id]
//   channel_id varchar [ref: > chat_channels.id]
//   created_at timestamp
//   updated_at timestamp
//   deleted_at timestamp [null]
//   type enum // text, image, link, 
//   body varchar
//   url varchar
// }


enum employment_types {
  full_time
  part_time
  intern
  contract
  freelance
}

enum company_size {
  startup
  small
  medium
  big
}
```
