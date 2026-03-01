# Settings

Settings is a unified panel accessed via the Settings link at the bottom of the left sidebar (below project navigation).

The panel has a left nav divided into three groups: Account, Workspace, and System.
The active project context carries over; switching projects from within Settings updates the Workspace sections accordingly.

---

## Navigation Structure

```
ACCOUNT
  Profile
  Notifications

WORKSPACE
  Project
  Team Members
  Logs

SYSTEM
  Admin         (Super Admin only)
```

---

## Account

### Profile
- Display name
- Email address
- Avatar / profile photo

### Notifications

Sets the default notification level for any ticket the user interacts with.
Per-ticket overrides remain available on each ticket itself.

Options:
- All activity (default)
- Mentions and Asks only
- QA progress only
- Status changes only
- None

---

## Workspace

Workspace sections are project-scoped. Visible only to users with Project Admin role on the active project.

### Project

- Project name
- Project mode: Personal / Delivery (toggle)

Mode transition rules:
- Personal to Delivery is allowed; confirmation required
- Delivery to Personal is restricted
- On transition to Delivery: existing tasks remain unchanged, no auto-conversion to tickets

Danger Zone (bottom of the page):
- Archive project (confirmation required)

### Team Members

- List of all project members with their current roles
- Role assignment per member: DEV / QA / PM / Project Admin (multi-select; a user can hold multiple roles)
- Invite user:
  - Search by email or username
  - If the account already exists: link the user to the project directly
  - If the account does not exist: send an email invite; account is created on acceptance
- Remove user from project

Role visibility note: tabs appear only when accessible. No disabled states or upgrade prompts.

### Logs

Audit log scoped to the active project. Visible to Project Admins (own project) and Super Admin (all projects).

Super Admin sees a project filter at the top. Project Admins see only their project's events.

What is logged:

| Event Type | Description |
|---|---|
| Role change | User, old role, new role, changed by, timestamp |
| Membership change | User, action (invited / joined / removed), by whom, timestamp |
| Destructive action | Action type (archive, soft delete), object, actor, timestamp |
| Debug access | Actor, object accessed, reason, issue reference, consent type, timestamp |

Properties:
- Immutable and append-only; no editing or deletion
- Filterable by: event type, date range, user
- Displayed in reverse chronological order

---

## System

### Admin

Visible to Super Admin only.

**Users**
- List of all accounts in the instance
- Add new account (email, display name)
- Deactivate / reactivate account
- View which projects each user belongs to (read-only)

**Debug Access**
- Lists active debug consent grants
- Each entry: user who granted consent, issue reference, granted at timestamp
- Super Admin can revoke at any time
- All accesses made under debug consent are automatically recorded in Logs

---

## What is NOT in Settings

- Per-ticket notification overrides (set on the ticket itself)
- Task privacy (tasks are always private by design, no toggle)
- Doc visibility (public/private is set per-doc at creation)
- Role-based tab visibility (automatic, not configurable)

---

## Visibility Summary

| Section | Who Sees It |
|---|---|
| Account > Profile | Every user |
| Account > Notifications | Every user |
| Workspace > Project | Project Admin |
| Workspace > Team Members | Project Admin |
| Workspace > Logs | Project Admin (own project); Super Admin (all projects) |
| System > Admin | Super Admin only |
