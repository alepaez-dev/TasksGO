# TasksGO


### License
Copyright (c) 2026 Alejandra Hernandez Paez

This project is licensed under the [PolyForm Noncommercial License 1.0.0](https://polyformproject.org/licenses/noncommercial/1.0.0/).

You may use this software for noncommercial purposes only.
Commercial use requires a separate license. Contact the author for inquiries.

### Description

TasksGO is a project focused on building a modern task and ticket
management system, designed to explore clean architecture and backend systems in Go.


## High-Level Concept (🔄)
TasksGO will support:
- Projects, which group work
- Tasks, which represent small actionable items, unrelated or related to tickets
- Roles: PM, DEV, QA. 🔄❓
- Tickets, which represent issues, features or requests
- Notes, attached to tickets (can change depending on the role)

### Tickets

A ticket should be a living workspace, not just a status tracker. It should be a place where decisions, context, and thinking evolve. Each role (DEV, QA, PM) interacts with tickets differently, so the view should adapt to how each role thinks.

**Core Features:**
- Planning notes (markdown)
- Visibility toggle (private to author / team visible)
- Attachments (later)

> Note: This is not a note app; note scope should remain small and ticket-focused.

#### DEV View
- Technical planning and implementation notes
- Private thinking space (debugging, exploration, trade-offs)
- Ability to tag notes for QA (e.g., `[QA]` tag auto-surfaces to QA view)

#### QA View
- Test scenarios and edge cases
- QA-focused summaries (generated from dev notes with `[QA]` tags)
- Bug reproduction steps (need to think how to do this cleanly, if I go back 2 months later to read ticket I need to understand what was tested)
- QA needs a standard structured way to mark scenarios as ***failing***, ***pending***, ***verified***. (scenarios are not comments, or notes.)-> how can I do this in the UI?

#### PM View (TBD, leaving to the end). Maybe PM view should be the standard.
- Status and progress summary
- Acceptance criteria
- Business context and stakeholder decisions
- Dependencies and blockers
- Timeline
- High-level approach (maybe)

**Example: DEV notes with QA tagging**
```
- Edge case: status being overridden
- Race condition if two people change the ticket at the same time
- Regression risk on flow
- [QA] Verify race condition doesn't happen when two people change the ticket simultaneously
```
The `[QA]` tag automatically surfaces that item to the QA view, saving dev time while keeping QA informed.

> Note: Dev thinking -> QA intent -> PM visibility

---

### Open Questions

**Important:**
- What makes us different from simply appending a Claude extension and asking, “What was tested here?” OR can we work with AI because the data is clearly shown in the ticket? !!!!!!!!!!!

**Less urgent:**
- Multiple devs/QAs working on the same ticket can get messy. How do we handle ownership and visibility when a ticket is poorly scoped by the PM?
- How can we help PMs create better tickets (AI?)? NTH, not the main business idea; looks like an AI wrapper. A good ticket creation assistant with good data might be useful, but getting good data is the important thing, not AI.

---

The exact boundaries and relationships between these concepts will be refined as the project evolves.

## Tech Stack (Initial) (🔄)
- Backend: Go
- Frontend: React or Next.js (TBD)
- AI: Separate service (design TBD) -- python
- Database: TBD