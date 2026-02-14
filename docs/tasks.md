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
- **private and personal** (only visible to the task owner, never to other users)

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

Tasks are **not**:
- tickets
- requirements
- verification artifacts
- documentation

If work requires shared ownership, verification, or evidence,
it should be promoted to a **ticket**.

---

## Tasks and Projects

- Every task belongs to a project
- Tasks created without an explicit project go to the default project
- Tasks exist in all project modes

In Delivery projects:
- Tasks support execution work
- Tickets represent delivery commitments

---

## Task State (v1)

Tasks intentionally have a **binary state only**:

- **Undone**
- **Done**

There are no additional states such as "in progress", "blocked", "waiting", or "converted".

Tasks can be **archived** to hide them from active views.
Archiving is a visibility action, not a state change.
A task remains Done or Undone even when archived.
Due date -> v2
---

## Creating Tasks

Tasks can be created in multiple ways:

- From the global Tasks view
- From a right side panel available anywhere
- From Dev Notes inside a ticket (task will be linked to the ticket)
- Only assigne possible is the user (PM can't assign to Dev, tasks are private and personal)

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
```

(this can be a config setting.)
Editing the `#task` line in Dev Notes syncs the task title.
The task title always reflects the current text of the originating line.

Deleting a note line never deletes the linked task.
The task persists independently once created.

---

## Completing Tasks

Tasks can be completed:
- from the global Tasks view (primary)
- from ticket context via a task side panel (secondary)

Tasks are never completed implicitly through text edits.

---

## Priority (v1)

Tasks can have a priority.

Rules:
- If a task has a custom priority, use it
- Else if linked to a ticket, use ticket priority
- Else, no priority (default)

Changing a ticket's priority updates all linked tasks that have not overridden it.

---

## Task and Ticket State Independence

Task state and ticket state are **completely independent**.

- Completing a ticket does not complete its linked tasks
- Completing a task does not affect the linked ticket
- A task can remain Undone after its linked ticket is Done

This is intentional. Tasks often outlive tickets:
- "When this ticket is deployed, update the staging config"
- "Write a retrospective note after release"


---

## Tasks and Docs

Tasks can be converted into Docs when execution turns into reusable knowledge. (not sure about this, makes sense for study projects, but seems like it is adding limited value, so it shouldn't be something we should focus on the UI interface) -> TODO: I need to think about this

Conversion:
- creates a Doc using the task's title and description
- links the Task ↔ Doc -> 0 value on the UI side, few value on the engineering side, maybe for reports, but for what else ?
- does not introduce a new task state

A task may be:
- done and linked to a doc
- undone and linked to a doc

"Converted to doc" is a relationship, not a status.

