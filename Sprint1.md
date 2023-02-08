# Sprint 1
## User Stories
1. As a site user, I want to be able to log in to the website to save my personal information.
2. As a site user, I would like to create financial goals so that I can be focused on my savings and spending to achieve this goal.
3. As a site user, I want to be able to easily add an expense so that my progress can be tracked.
4. As a site user, I want to be able to easily add a planned income so that my progress can be tracked.
5. As a site user, I want my log in to be more secure through two-factor authentication so that people are less likely to steal my information.
6. As someone new to budgeting, I want to be guided through the features so that I can start using the app with ease.
7. As an officer of a club, I want to be able to invite members of my team, so we can track our spending together.
8. As a site user, I would like to connect my bank account to the site to transfer money based on my financial goals.
9. As a frequent user, I want to open my budget without having to log in each time, so I can access my budget tracker faster.
10. As a site user, I want to see my savings progress visually to easily analyze my financial data and trends.
11. As a site user, I would like to have my financial data displayed in an organized fashion, so I won't be overwhelmed.

## Issues We Planned to Address
We wanted to focus on the login for this first sprint. Financial data is important and needs to be kept safe. Because of that, we wanted to focus on story 1. Our goal was to create a system that allowed the user to sign up with their name, email, and password, then sign in with their email and password. The user would then be brought to a dashboard page, displaying their data. Entering incorrect credentials or the credentials of a non-registered user would prevent the user from logging in.

## Issues We Successfully Completed
We were able to create a login page for registered users which covers story 1. The user is able to register using the sign-up page. Their password is hashed and then stored along with their email in a database. The login page allows the user to enter their credentials and login to the system. If successful, the backend will create a cookie in the browser containing a JSON web token. The user is then brought to the dashboard page where the frontend will check for the token, and then request the user's data from the backend. Their name is then displayed onscreen.

## Issues Not Successfuly Completed
We have not fully implemented a logout system yet. When the user is logged in, they will remain logged in until someone else logs in, or the cookie expires. We needed to spend more time on figuring out how to connect the frontend with the backend. We plan to implement the logout features in the upcoming sprint. Additionally, our login system lacks input validation, which can introduce security vulnerabilities. We wanted to approach our login system one step at a time and focus strictly on getting the system working first. We could step up the security measures on the login system by creating an input validation method in a future sprint.
