-- +goose Up
-- +goose StatementBegin
create table urls_protected
(
    id    serial primary key,
    url   text not null,
    roles json not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table urls_protected;
-- +goose StatementEnd
