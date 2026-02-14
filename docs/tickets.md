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
- Anyone can create a ticket (any role)

> This is not a general note app.  
> Notes must remain **ticket-scoped and intentional**.

---

# Ticket Views (Role-Based Tabs) 🔸

Each ticket is divided into **role-based views** implemented as tabs.

See [Users & Roles](users.md) for detailed role permissions.

[ Overview ] [ Dev ] [ QA ] [ Activity ]

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

**Privacy**: Only the dev assigned to the ticket can see their own dev notes.
Other devs on the project cannot access them.

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

`[QA]` tagging is the **only mechanism** that makes Dev Notes content visible to others.

When a dev writes a `[QA]` tag on a line, they are **explicitly choosing** to publish that line to the QA view as a scenario. All other Dev Notes content remains strictly private to the dev.

This is an intentional act, not an automatic extraction. The dev controls what crosses the privacy boundary.

Dev thinking (private) → `[QA]` tag (explicit publish) → QA scenario (shared) → PM visibility via QA summary

---

## QA View 👁️ (Verification Space)

Purpose: explicit testing and accountability.

QA scenarios are not comments and not notes.
They are verifiable assertions with state.

Evidence over optimism.

When QA verifies an scenario or a ticket they need to add evidence. (text forcefully and images/videos optional)

---

# Scenarios 🔸

## Concept

A scenario represents a concrete behavior or edge case that must be verified.

Scenarios:
- Are finite
- Have ownership
- Have explicit state
- Carry their own evidence

See [Users & Roles](users.md) for scenario permissions.

QA owns the final checklist. This prevents silent scope reduction and creates clear accountability.

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
A toggle with "Create tasks for all upcoming scenarios"

---

# Open Questions (V1)
1. What makes us different from simply appending a Claude extension and asking, "What was tested here?" OR can we work with AI because the data is clearly shown in the ticket? !!!!!!!!!!!
2. Multiple devs/QAs working on the same ticket can get messy. How do we handle ownership and visibility when a ticket is poorly scoped by the PM?
3. How can we help PMs create better tickets (AI?)? NTH, not the main business idea; looks like an AI wrapper. A good ticket creation assistant with good data might be useful, but getting good data is the important thing, not AI.

# AI Features (V2)

Not core for v1. v1 focus is collecting good structured data. AI amplifies that data later.

## Planned Features
- Ticket assistant creator (vector DB for similar tickets, topics, docs)
- Ticket assistant reports (requires reports feature first)
- "What scenarios have failed before on similar tickets?"

## Ticket Intents

Intents are dynamic topics extracted from tickets (e.g., "task ownership", "email notifications").
Projects have different domains; intents should emerge from data, not be predefined.

**Phase 1: Similarity-based**
Tickets are converted to vectors embeddings. When creating a ticket, show related tickets based on vector proximity.

**Phase 2: Clustering-based** (overengineering if we don't have the data)
Once a project has enough tickets (20+) do clustering to group tickets into named topics. Enables browsing by category
