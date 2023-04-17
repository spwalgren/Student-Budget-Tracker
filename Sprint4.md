# Work Completed in Sprint 4

**Main Points:**

# Frontend Unit Tests

For unit testing, we opted to use Cypress.
The code for writing component tests in Cypress is less complex than Jasmine tests and is easier to understand.
Additionally, it offers useful tools such as running the tests in a browser and reviewing how Cypress sees the individual components.
Seeing that the Cypress component tests accomplish the same tasks as the Jasmine tests, we believed it to be an acceptable substitute.

## New/Changed for Sprint 4

**components/dash-budgets** - Tests if it can add tables, add to existing tables, delete data from tables, delete tables, and edit table entries.
Changed to be compatible with the current version that uses an API to calculate period information.

**components/dash-transactions** - Tests if it has a table, has a button that opens a modal, can display entries from the mock service, has a detail row that pops out, can add transactions, can edit transactions, and can delete transactions.
Changed to test the component more thoroughly with add, edit, and delete functions.

**components/event-card** - Tests how it displays information in each of its states, including "On Track", "Upcoming", "Overspent", and "Completed".

**components/dash-calendar** - Tests if it can display the current month and can navigate between months.

## Unchanged from Sprint 3

**components/alert** - Tests how it displays colors, text, and icons.

**components/dash-home** - Tests if it can mount

**components/dash-settings** - Tests if the delete user button behaves as it should and calls a function.

**components/page-not-found** - Tests if it has a button and if that button calls a function when clicked.

**routes/landing** - Tests if it can mount

**routes/dashboard** - Tests if it has visible buttons, it calls a function immediately, and it calls a function upon clicking log out.

**routes/login** - Tests if it automatically attempts to log in and tests various inputs to see what will permit the user to log in.

**routes/signup** - Tests various inputs to see what will allow the user to log in.