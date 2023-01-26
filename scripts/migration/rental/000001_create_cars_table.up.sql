create table if not exists cars
(
    id           varchar(10)                         not null,
    area         text      default ''                not null,
    project_id   text      default ''                not null,
    project_name text      default ''                not null,
    seat         tinyint   default 0                 not null,
    type_name    text      default ''                not null,
    latitude     float     default 0                 not null,
    longitude    float     default 0                 not null,
    status       int       default 0                 not null,
    created_at   timestamp default CURRENT_TIMESTAMP not null,
    updated_at   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint pk_cars primary key (id)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment = 'cars';

create index idx_status on cars (status);
