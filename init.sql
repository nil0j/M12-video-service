/* run: */
create database jirafeitor;
create user jirafeitor_user with password 'changeme';
\connect jirafeitor;
grant all privileges on database jirafeitor to jirafeitor_user;
