# Tasks

Tasks represent **execution** in TasksGO.

They answer a simple question:

> “What am I actually going to do?”

Tasks are lightweight, actionable, and always available,
regardless of project mode.

---

## What Tasks Are

Tasks are:
- actionable
- meant to be completed
- PERSONAL

Examples:
- Investigate webhook race condition
- Refactor auth middleware
- Study Go interfaces
- Review PR

Tasks do **not** require:
- QA
- PM involvement
- evidence
- delivery readiness

---

## What Tasks Are Not

Tasks are not:
- tickets
- requirements
- verification artifacts
- documentation

If work requires shared ownership, verification, or evidence,
it should be promoted to a **ticket**.

---

## Tasks and Projects

- Every task belongs to a project
- Tasks exist in all project modes
- Tasks are never disabled

In Delivery projects:
- Tasks support execution work
- Tickets represent delivery commitments

---

## Creating Tasks

Tasks can be created in multiple ways:

- From the global Tasks view
- From a side panel available anywhere
- From Dev Notes inside a ticket

Creation should be fast

---

## Tasks Created from Dev Notes

Inside a ticket’s Dev Notes, tasks can be created using a simple tag:

```md
- Investigate webhook race condition #task
```
(this is gonna be interesting to achieve.)


When #task is used:
- A task is created immediately
- The task title comes from the line text
- The task is linked to the ticket
- No modal or interruption occurs

Descriptions are optional.

Indented lines can be used to add context:
```md
- Investigate webhook race condition #task
  - Happens when webhook arrives before DB commit
  - Suspect transaction boundary issue
```md