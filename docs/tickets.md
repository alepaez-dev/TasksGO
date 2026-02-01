# Tickets 🔸

**Built for thinking, not ticket shuffling.**

A ticket is a **living workspace**, not just a status tracker.
It is the place where context, decisions, implementation, and verification
evolve together.

TasksGO intentionally ignores Scrum and Agile ceremonies.
Instead, it focuses on how work actually happens.

> **Clarity and evidence over rituals.**

Each role (PM, Dev, QA) interacts with a ticket differently, so the ticket adapts
its structure and visibility to how each role thinks and works.

---

## Core Principles

- Tickets are workspaces, not comment threads
- Thinking is allowed to be messy
- Verification is explicit and structured
- Ownership is clear
- Noise is minimized by design

> This is not a general note app.  
> Notes must remain **ticket-scoped and intentional**.

---

# Ticket Views (Role-Based Tabs) 🔸

Each ticket is divided into **role-based views** implemented as tabs.

If a role is not allowed to see a view, **the tab does not appear**.

[ Overview ] [ Dev ] [ QA ] [ Activity ]
| Role | Overview | Dev | QA | Activity |
|------|----------|-----|----|----------|
| PM   | ✅        | ❌  | ❌ | ✅       |
| Dev  | ✅        | ✅  | ❌ | ✅       |
| QA   | ✅        | ❌  | ✅ | ✅       |
PMs do not access Dev or QA tabs by default.
Instead, they see **summaries and shortcuts** in the Overview.

--- 

## Overview View 👁️ (PM-First)
Purpose: **“Can we ship this?”**
The Overview answers this question without requiring PMs to read
implementation or test details.

### Contents
- Status and readiness indicator
- Description (what)
- Business context (why)
- Scope
- Dependencies and blockers
- Timeline (optional)
- QA summary (expandable, read-only)

Example:

```md
### Description
Users are not receiving confirmation emails after tour scheduling.

### Why
This blocks leasing conversions and causes manual follow-ups.

### Scope
Fix confirmation email logic for scheduled tours only.
```
QA Summary (PM-Friendly)
- 3 / 5 scenarios verified
- 1 failing
- Last verified 2h ago

PMs do not see test evidence or attachments here.
Evidence lives at the scenario level to prevent attachment dumps.

---
## Dev View 👁️ (Thinking Space)
Purpose: engineering thinking, not documentation.

### Dev Notes
- Markdown
- Editable
- Allowed to be incomplete or incorrect
- Not required to be cleaned up
- May be private or team-visible

Dev notes are temporary thinking artifacts:
they support reasoning, exploration, and decision-making,
not final documentation.
Example:
```md
- Possible race condition if two users update simultaneously
- State recalculation might be overridden
- [QA] Verify race condition does not occur with concurrent updates
```
### QA Tagging

Lines tagged with [QA] automatically surface in the QA view as an scenario.

Dev thinking → QA intent → PM visibility

---

## QA View 👁️ (Verification Space)

Purpose: explicit testing and accountability.

QA scenarios are not comments and not notes.
They are verifiable assertions with state.

Evidence over optimism.

---

# Scenarios 🔸

## Concept

A scenario represents a concrete behavior or edge case that must be verified.

Scenarios:
- Are finite
- Have ownership
- Have explicit state
- Carry their own evidence

Permissions
- Dev can create scenarios
- QA can create, edit, archive, or delete scenarios
- PM cannot create scenarios (v1)

Once created:
- Dev cannot delete scenarios
- Dev cannot mark scenarios out of scope (this will need an interaction in the comments)
- QA owns the final checklist

This prevents silent scope reduction and creates clear accountability.

---

## Scenario States (v1)
- Not Tested
- Failing
- Fixed — Needs Retest
- Verified
- (Optional) Blocked
- (Optional) Out of Scope

--- 

## Evidence Handling
Evidence is attached per scenario, not per ticket
- Supports images, videos, logs, and links
- Collapsed by default
- Displayed via an indicator: 📎 2

When a scenario is Verified, it collapses to a single line:
```
[✓] Scenario title — Verified by QAName · Jan 31, 2026 · 📎
```

If a verified scenario is reverted:
- New evidence is required
- Old evidence cannot be reused

This ensures verification remains trustworthy over time.

---

# Activity Tab 💬

Purpose: shared discussion without noise.
- All comments live here
- Comments are anchored to ticket sections
- Clicking a comment jumps to its context
- No comments live inside Dev or QA tabs

Hovering over major sections shows a 💬 icon to ask contextual questions.

---

Notifications (Per Ticket)
- All activity
- Mentions / questions only
- QA progress
- Status changes

Configurable per ticket (default: All).

---

# Open Questions (V1)
1. What makes us different from simply appending a Claude extension and asking, "What was tested here?" OR can we work with AI because the data is clearly shown in the ticket? !!!!!!!!!!!
2. Multiple devs/QAs working on the same ticket can get messy. How do we handle ownership and visibility when a ticket is poorly scoped by the PM?
3. How can we help PMs create better tickets (AI?)? NTH, not the main business idea; looks like an AI wrapper. A good ticket creation assistant with good data might be useful, but getting good data is the important thing, not AI.