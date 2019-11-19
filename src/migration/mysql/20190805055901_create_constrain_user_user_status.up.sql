ALTER TABLE `users` 
ADD CONSTRAINT `users_user_status` 
FOREIGN KEY(user_status_id) 
REFERENCES user_status(id)