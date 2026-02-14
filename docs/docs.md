# Docs

Docs represent **knowledge** in TasksGO.

They are used for information that should outlive a task or ticket.

Docs are intentionally simple.

> Tasks are for doing.  
> Docs are for remembering.

---

## What Docs Are

Docs are:
- lightweight markdown pages
- project-scoped
- meant for reference and learning

Examples:
- Study notes (e.g. “Go Interfaces”)
- Design decisions
- Troubleshooting guides
- Documentation

---

## What Docs Are Not

Docs are not:
- a full notes app
- a replacement for Dev Notes
- a workflow tool
- hierarchical knowledge bases

There are no:
- nested documents
- complex formatting tools
- task-like states

---

## Docs and Projects

- Every doc belongs to a project
- Docs exist in all project modes

Docs do not affect:
- ticket status
- task state
- delivery readiness

---

## Doc Visibility and Permissions

Docs can be **public** or **private**.

- **Public**: visible to all project members
- **Private**: visible only to the doc creator

Docs also have **permission levels**:

- **Read**: can view the doc content
- **Edit**: can modify the doc content

Permission rules:
- The doc creator always has edit access
- Public docs are read-only to other project members by default
- The creator can grant edit access to specific users or roles
- Private docs are invisible to everyone except the creator

Default behavior:
- Docs created in Personal mode projects are private by default
- Docs created in Delivery mode projects are public (read-only) by default

---

## Creating Docs

Docs can be created:
- directly from the Docs view
- by converting a Task into a Doc


---

## Converting Tasks to Docs

A task can be converted into a doc when execution turns into knowledge.

When a task is converted:
- A new Doc is created
- Doc title = task title
- Doc body = task description (if any)
- Task and Doc are linked

The task state remains simple:
- usually marked as **Done**
- optionally archived

Conversion does not introduce a new task state.

---

## Docs vs Dev Notes

Dev Notes:
- live inside tickets
- are for temporary thinking
- evolve during problem-solving

Docs:
- live at the project level
- are meant to be read later
- capture stable understanding

Rule of thumb:
- “I’m figuring this out” → Dev Notes
- “I’ll want this later” → Doc

---

## Design Principles

- Docs stay simple
- Docs are flat (no hierarchy)
- Docs do not compete with tasks or tickets
- Visibility is explicit: public or private
- The creator controls access