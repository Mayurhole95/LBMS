CREATE TABLE `Transaction`(
    `id` VARCHAR(40) NOT NULL,
    `issuedate` VARCHAR(100) NOT NULL,
    `returndate` VARCHAR(100) ,
    `duedate` VARCHAR(100) NOT NULL,
    `book_id` VARCHAR(40) NOT NULL,
    `user_id` VARCHAR(40) NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES user(id)
    on delete cascade,
    FOREIGN KEY(book_id) REFERENCES book(id)
    on delete cascade
    );