# Users & Roles (v1)

This document defines how users, roles, project administration, and privacy work in TasksGO.

TasksGO prioritizes:
- clarity of responsibility
- calm, role-based visibility
- user trust via explicit consent and audit trails

---

## Core Concepts

### User
A **User** is a person with an account in TasksGO.

Users can belong to multiple projects and can have different responsibilities per project.

### Roles are project-scoped
Roles are assigned **per project**, not globally.

A user may hold multiple roles in the same project (e.g. solo builder):
- DEV + QA + PM

### Visibility over permission toggles
If a user cannot access a view, the tab does not appear.
No disabled buttons or "upgrade to access" teasers.

---

## Global Role

### Super Admin (Instance Owner)
A **Super Admin** is the owner of the TasksGO instance (e.g. the founder / operator).

Super Admin powers exist for:
- account recovery
- admin operations
- debugging assistance

Super Admin does not automatically access private user content.
Private access requires **explicit debug consent** (see Debug Access).

---

## Project Roles

### Project Admin
A **Project Admin** manages project membership and project roles.

Project Admin can:
- invite/remove users from the project
- assign/remove project roles (DEV / QA / PM)
- manage project settings (name, visibility)
- (optional) change project mode (Personal ↔ Delivery)

Project Admin does **not** automatically see private notes.

---

### DEV (Developer)
Purpose:
- technical ownership and implementation

Can:
- create tickets (Delivery mode only)
- view ticket Overview + Activity + QA tab (read-only)
- access Dev tab (only for tickets they are assigned to)
- write and manage Dev Notes
- create tasks (global + from Dev Notes)
- create QA scenarios (but not verify them)

Cannot:
- edit QA scenarios (unless also QA)
- mark scenarios as Verified (unless also QA)
- delete QA scenarios (unless also QA)
- manage scenario evidence (unless also QA)
- see Dev Notes from other devs (even on the same project)

**Dev Notes Privacy**: A dev can only see their own Dev Notes on tickets they are assigned to. Other devs cannot access them.

---

### QA (Quality Assurance)
Purpose:
- verification ownership and evidence

Can:
- create tickets (Delivery mode only)
- view ticket Overview + Activity + QA tab
- edit QA tab (only role with edit access)
- create/edit/delete QA scenarios (soft delete)
- mark scenarios as Pending / Failing / Verified
- attach and manage scenario evidence

Cannot:
- access Dev tab (unless also DEV)
- edit Dev Notes (unless also DEV)

---

### PM (Product Manager)
Purpose:
- business intent, delivery clarity, readiness understanding

Can:
- create tickets (Delivery mode only)
- view ticket Overview + Activity + QA tab (read-only)
- view QA Summary in Overview
- comment and ask questions

Cannot (v1):
- access Dev tab
- edit QA tab
- create/edit QA scenarios
- edit Dev Notes

---

## Project Modes and Roles

### Personal Mode
- Tickets do not exist
- Tasks and Docs are available
- Roles are effectively lightweight (Contributor-like)

### Delivery Mode
- Tickets exist
- Overview is visible to project members
- Dev tab requires DEV role
- QA tab is visible to all, editable by QA only

Roles become relevant only when the project mode requires them.

---

## Ticket View Visibility (Summary)

In Delivery mode:

- Overview: visible to all project members
- Dev: visible only to assigned DEV
- QA: visible to all, editable by QA only
- Activity: visible to all project members

---

## Debug Access (Privacy & Trust)

TasksGO supports private content (e.g. private Dev Notes).
Private content must feel safe.

### Debug Consent
Users may report issues via support channels (email/SMS/etc).

When reporting an issue, a user can explicitly grant debug access via a checkbox:

> “I allow the TasksGO owner to temporarily access my private data for debugging purposes only related to this issue.”

Debug consent should be:
- explicit
- issue-scoped (not permanent)
- purpose-limited

### What debug access allows
With debug consent, Super Admin may access otherwise-private content strictly to investigate the reported issue.

### Audit Logging (Required)
Any debug access must be **audited**.

Audited means:
- the system records an immutable event describing the access

Minimum audit fields:
- who accessed (actor)
- what was accessed (object)
- when it happened (timestamp)
- why (reason / issue reference)
- how (debug access with user consent)

Audit logs exist to provide transparency and accountability.

---

## Audit Log (v1)

Audit logging should cover:
- debug access to private data
- role changes
- project membership changes
- destructive actions (delete/archive where relevant)

Audit logs are readable by:
- Super Admin
- Project Admin (for their project)

---

## Design Principles

- Roles are contextual to projects, not global identity, except for SuperAdmin
- Users can hold multiple roles in the same project
- Tabs appear only when accessible
- Private content requires explicit debug consent for Super Admin access
- Debug access is always audited