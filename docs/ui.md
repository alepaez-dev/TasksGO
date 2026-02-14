## Mockups
- Main layout - Done.
- Home page (notifications, activities, etc) - new|unsure -> we can show waiting on tickets, so when something is blocked is not lost on software notifications or emails. We can always check it in home page.
- Tasks page - ALL mockups done ✅
- Tickets - In progress
    - Overview tab mockups - ALL mockups Done ✅
    - Dev tab - ALL mockups done ✅ -> need to add same right aside component from Tasks page
    - QA - in progress
    - Activity - TBD.
- Docs - TBD


## Home v1
Inbox / Dashboard. **This is the default landing page.**

Purpose:

Tell the user what matters right now, without asking them to think.


It answers:

“What should I personally look at right now?”

Home page contains:
- Needs my attention
- Assigned to me
- Watching
- Activity

These are:
- personal
- cross-project (if we only have a project selected ignore)
- role-agnostic
- attention-based

Home shows:
- Needs my attention → Open Asks where the user hasn’t responded yet
- Assigned to me → Tickets the user is responsible for executing
- Watching → Tickets the user interacted with and is following
- Activity → Recent tickets the user touched (history)

---


## Tickets v1 — Filters, Flow, and Communication

### Tickets List — Filters (v1)

Filters answer what you should look at, not who is at fault.

Involvement (primary filter)
- Assigned to me
- Waiting on me
- Created by me
- Watching (optional)

Status
- Draft
- In Progress
- In Review
- Blocked
- Verified
- Done

Priority (optional)
- High
- Medium
- Low

Important: Filters are relationship-based, not role- or person-based.

---

### Inside a Ticket — Flow and Communication

The Overview is the shared context of a ticket.
It answers what, why, what is included, and where we stand — without noise.

Contents
- Status & Priority
- Current progress and urgency
- Description (What)
- What is being built or changed
- Why
- Business or technical reason
- Why this work matters
- QA Summary
- High-level testing state
- Scenarios verified / failed / waived
- Expandable, read-only
- Scope
- What is included
- What is explicitly excluded
- Metadata (side panel)
- Assignee, reporter, sprint, component, environments

Key Principles
- Overview is readable in under 30 seconds
- No implementation details
- No discussion threads
- No attachments dump

If something requires deep discussion, evidence, or execution:
→ it belongs in Dev, QA, or Activity, not here.



#### Ask vs Comment
There are two ways to write inside a ticket.

***Comment (default)***
Comments are used for discussion, updates, and context.
- Free-form communication
- Mentions allowed
- Sends notifications
- Never blocks progress
- Does not affect “Waiting on”

Comments are the default action in the UI and have keyboard shortcuts.
```@qa Ticket is ready, please test when you can```

----

***Ask***

Asks are used when input is needed to move forward.
- Explicit action chosen by the author
- Targets one or more specific people
- Creates a blocking object
- Affects “Waiting on”
- Visible and distinct in the UI
- Disabled until at least one person is targeted (tooltip shown)

```@ale What is the scope of this ticket?```

----

Ask Behavior

Per-target response tracking

When an Ask targets multiple people, each target is tracked individually.
- Once a person responds:
- The ticket is no longer waiting on them
- The Ask remains open if other targets have not responded

Ask resolution

An Ask is considered fully resolved when:
- All targeted people have responded, OR
- Someone explicitly clicks Resolve Ask

Resolving an Ask:
- Clears all remaining “waiting on” signals
- Stops notifications for that Ask
- Does not require a comment
- Is recorded in activity

This allows Asks to be resolved when:
- One answer is sufficient
- Agreement is implicit
- The answer happened outside the system

----

Waiting On

"Waiting on" represents a **user-level** blocking condition.

It answers:

Who are we still waiting on to move forward?

**Waiting On applies to people, not to other tickets.**
For ticket-to-ticket relationships, see "Depends On" below.

How "Waiting on" is computed

Waiting on is:
- Automatic
- Never manually set
- Derived from:
- Open Asks (remaining unanswered targets)
- Unresolved QA scenarios

If multiple blockers exist, only the most relevant waiting signal is shown.

----

Waiting on Me

A ticket is "waiting on me" if:

There exists at least one open Ask where I am a target and I have not yet responded.

Once I respond:
- The ticket immediately stops being "waiting on me"
- Even if the Ask remains open for others
- Even if the ticket is still blocked overall

----

Depends On (Ticket-to-Ticket Dependencies)

"Depends On" represents a **ticket-level** blocking condition.

It answers:

Which other tickets must be resolved before this one can proceed?

Rules:
- Dependencies are set manually by the ticket creator or assignee
- A ticket can depend on one or more other tickets
- Dependencies are visible in the Overview under "Dependencies and blockers"
- Depends On is separate from Waiting On; they serve different purposes
- Waiting On = "which person do I need a response from?"
- Depends On = "which other work must finish first?"
