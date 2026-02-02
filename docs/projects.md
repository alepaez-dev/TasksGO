# Projects

Projects are the top-level containers in TasksGO.
They represent **areas of work**, not workflows.

---

## Navigation

The UI has an always-visible left sidebar for navigation:

- **Tasks** (home)
- Tickets
- Docs
- Projects

Tasks is the default landing screen.

```
┌─────────────────────────────────────────────┐
│  [Logo]                                     │
├──────────┬──────────────────────────────────┤
│Project   │                                  │
│  Tasks ← │   Main content area              │
│  Tickets │                                  │
│  Docs    │                                  │
│          │                                  │
│          │                                  │
└──────────┴──────────────────────────────────┘
```

A project can be:
- a software system
- a feature area (TBD)
- a learning goal
- a personal initiative

Projects organize **tasks and tickets and docs**, but they do not dictate how work must be done.

Every user has a **default project**. Tasks created without an explicit project go there.

> Projects evolve. Tools should support that evolution without forcing process too early.

---

## Project Modes

Each project has a **mode** that defines the level of rigor required.

### 1. Personal Mode

Personal mode is lightweight and flexible.
It is intended for:
- learning
- exploration
- solo work
- early planning phases

In Personal mode:
- Tasks are always available
- Tickets are disabled
- No PM or QA concepts exist
- Focus is on execution and progress

Examples:
- “Study Go”
- “Learn System Design”
- “Personal Admin”

---

### 2. Delivery Mode

Delivery mode is for work that requires:
- coordination
- shared ownership
- verification
- evidence of correctness

In Delivery mode:
- Tasks are available
- Tickets are enabled
- QA workflows exist
- PM-focused summaries and readiness indicators exist

Examples:
- “TasksGO Backend”
- “Payments Refactor”
- “Release v1 API”

---

## Project Mode Transitions

Projects can transition from **Personal → Delivery**.

This transition is:
- intentional

When a project moves to Delivery mode:
- Existing tasks remain unchanged
- No tasks are automatically converted into tickets
- Tickets become available

The reverse transition (Delivery → Personal) is restricted.

---

## Tasks vs Tickets in Projects

Projects **always support tasks**.

- Tasks represent execution and personal commitments
- Tickets represent delivery commitments that require verification

Even in Delivery mode:
- Not all work needs to be a ticket
- Tasks remain the default tool for day-to-day execution

---

## Linking Tasks and Tickets

- Tasks can optionally link to tickets
- Tickets can reference related tasks
- Links are lightweight and non-intrusive

Tasks never live inside tickets by default.
They remain accessible from the global Tasks view.

---

## Design Principles

- Projects define **context**, not process
- Execution (tasks) and delivery (tickets) remain clearly separated