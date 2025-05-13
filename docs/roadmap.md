#  Roadmap: Solo Adventure Picker

This roadmap outlines the development phases for turning the app into a gamified real-world RPG experience. It breaks features into logical chunks from MVP to mobile readiness.

---

##  Phase 0 – Foundation (COMPLETE)
- [x] Set up Go backend with routes and HTTP server
- [x] Set up Vue frontend and basic component structure
- [x] Connect frontend to backend with fetch
- [x] Style base page with CSS
- [x] Successfully reroll random adventure via button

---

##  Phase 1 – MVP Loop
- [ ] Design and apply animated card transition when rerolling
- [ ] Display XP value for adventure types
- [ ] Implement XP counter per user (mock data first)
- [ ] Create journal input box (with basic validation)
- [ ] Award bonus XP on journal submission
- [ ] Create and store user ID in localStorage (pre-auth)
- [ ] Add adventure filters (type, mood, etc.)
- [ ] Style page for mobile responsiveness
- [ ] Log daily session in `SESSION_LOG.md`

---

##  Phase 2 – Gamification Layer
- [ ] Add user levels based on XP
- [ ] Animate level-up (popup or card glow)
- [ ] Create and display achievements (e.g., 5 hikes, 1 journal entry)
- [ ] Add reroll limit (reset daily)
- [ ] Reward reroll tokens from achievements or journaling
- [ ] Daily quest system (e.g., “Do 1 new thing”)
- [ ] Add `/docs/xp-system.md` and define XP formula
- [ ] Style gear badges or perks (no functionality yet)

---

##  Phase 3 – World Map & Fog of War
- [ ] Create map region system (tile/grid or real map)
- [ ] Apply fog overlay and track which tiles are cleared
- [ ] Link adventures to map locations
- [ ] Mark visited adventures on the map
- [ ] Animate fog clearing
- [ ] Show user’s current region/progress
- [ ] Add `/docs/fog-of-war.md` to explain logic + UX plan

---

##  Phase 4 – Persistent User System
- [ ] Add proper user auth (Supabase or custom JWT)
- [ ] Store visited adventures and XP in MongoDB
- [ ] Load journal entries from DB
- [ ] Let users view past adventures
- [ ] Add settings page for account and preferences

---

##  Phase 5 – Mobile & Real-World Integration
- [ ] Convert app into mobile-friendly PWA or native wrapper
- [ ] Use GPS to suggest nearby adventures
- [ ] Add location-based fog reveal
- [ ] Estimate travel time and duration of adventure
- [ ] Add optional cost estimator (admission, food, gas)
- [ ] Add stretch rewards for multi-day streaks

---

##  Stretch Goals / Long-Term
- Rare gear unlocks (cosmetics or perks)
- Multiple character “classes” (Explorer, Foodie, etc.)
- Guild system to join groups
- Public map sharing with journal visibility
- Monetization: cosmetics or reroll token packs
