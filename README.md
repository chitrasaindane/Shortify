<!-- # "Shortify" Readme File # -->
# 🕸️ Welcome to Shortify - Modern URL Shortener 🕸️

Shortify is a modern, full‑stack URL shortening service built with Go (backend), MongoDB (data store), Clerk (authentication), and Next.js (frontend). It's designed for rapid development and production readiness with clean APIs and a polished UI.

---

## 🚀 Live Deployments

- 🌐 **Frontend (Vercel)**  
  👉 https://myshortify.vercel.app

- ⚙️ **Backend (Render)**  
  👉 https://shortwave.onrender.com

---

## Tech Stack

- **Backend:** Go
- **Database:** MongoDB
- **Auth:** Clerk
- **Frontend:** Next.js

---

## Key Features

- Shorten URLs into vanity links (username + slug)
- Fast redirect service for shortened links
- User management and authentication via Clerk
- MongoDB for reliable storage of users and links
- Next.js frontend with responsive UI and Clerk integration

---

## Project Structure (high level)

- `Server/` — Go backend, API handlers, app logic and MongoDB integration
- `Client/` — Next.js frontend, components, and Clerk auth routes
- `README.md` — This file

---

## Quick Start

Prerequisites:

- Go 1.20+ installed
- MongoDB accessible (local or hosted)
- Clerk account and API keys (for auth)

1. Clone the repo

```bash
mkdir Shortify
cd Shortify
git clone https://github.com/<your-github-username>/Shortify.git .
```

2. Configure application settings (local & production)

Before running the app locally or deploying, make sure your runtime configuration is in place:

- Local development: copy `Server/.env.example` to `Server/.env` and `Client/.env.example` to `Client/.env`, then open those files and fill in values specific to your environment. Keep sensitive values out of version control.

- Production: configure required settings and secrets in your hosting platform's environment manager (Vercel, Render, etc.). Ensure callback/webhook URLs (for Clerk and other services) point to your deployed URLs.

Use the example files in `Server/.env.example` and `Client/.env.example` as the authoritative list of keys to provide; the README avoids enumerating individual variable names to keep configuration details centralized in the example files.

3. Run the backend

```bash
cd Server
go mod tidy
go build main.go
./main
```

4. Frontend: install and run

```bash
cd Client
npm install
npm run build 
npm run dev
```

The frontend typically runs on `http://localhost:3000` and the backend on `http://localhost:3001` (configurable).

---

## Common Endpoints

- Redirect (public): `GET /{username}/{slug}` — redirects to the destination URL of the link
- API base: `/api/v1`
  - User webhook: `POST /api/v1/user/webhook` (Clerk)
  - Update username: `PUT /api/v1/user/username`
  - Links CRUD: under `/api/v1/link`

---

## Testing & Development Tips

- Use `curl -v http://localhost:3001/username/slug` to test redirects.
- Run `go build main.go` to check for backend compile errors.
- Ensure Clerk webhooks point to the server's `/api/v1/user/webhook` during integration.

---

## Deployment

Below are quick deployment flows for the frontend (Vercel) and backend (Render). These are minimal steps — adapt them for your environment and secrets manager.

- Vercel (Frontend)

  1. Create a Vercel project and connect it to this repository.
  2. Set the Project Root to `Client` (or import as a monorepo and point the app to `Client`).
  3. Build command: `npm run build`
  4. Output directory: leave default (Next.js handled by Vercel).
  5. Configure required settings and secrets in Vercel (Dashboard → Settings → Environment Variables).
  6. Deploy — Vercel will run builds on every push.

- Render (Backend)

  1. Create a new Web Service on Render and connect it to the repository.
  2. Set the root directory to `Server`.
  3. Environment & Build:
    - Build command: `go build main.go`
    - Start command: `./main`
  4. Configure required settings and secrets in Render (Service settings → Environment).
  5. Deploy — Render will build and start the service; check logs for startup errors.

Tips:

- Use Vercel for the Next.js frontend (serverless/edge‑optimized) and Render (or similar) for the Go backend.
- Keep production secrets in the platform's environment manager — never commit them.
- If using webhooks (Clerk), configure callback URLs in Clerk to point to your deployed `POST /api/v1/user/webhook` endpoint.

---

## Contributing

Contributions welcome — open issues or submit PRs. Follow these steps:

1. Fork the repo
2. Create a feature branch
3. Make changes and run `go build` and `npm run dev` locally
4. Open a PR with a clear description

---

## License

This project uses the license in the repository. Feel free to adapt as needed.

---

## Acknowledgments
Enjoy Shortify — deployment and CI-CD guidance have been added above. Open an issue or PR if you want badges, examples, or more detailed deployment templates.

---