# Step 1. Rebuild the source code only when needed
FROM node:20-alpine AS builder

WORKDIR /app

# Install dependencies based on the preferred package manager
COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml* ./
# Omit --production flag for TypeScript devDependencies
RUN \
  if [ -f yarn.lock ]; then yarn --frozen-lockfile; \
  elif [ -f package-lock.json ]; then npm ci; \
  elif [ -f pnpm-lock.yaml ]; then yarn global add pnpm && pnpm i; \
  else echo "Lockfile not found." && exit 1; \
  fi
COPY . .
# Environment variables must be present at build time
# https://github.com/vercel/next.js/discussions/14030

# This will do the trick, use the corresponding env file for each environment.
COPY .env.production .env.production
ENV NEXT_TELEMETRY_DISABLED 1
RUN yarn build

# Step 2. Production image, copy all the files and run next
FROM node:20-alpine AS runner

WORKDIR /app

# Don't run production as root
RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

# sharp is needed for standalone mode
RUN npm i sharp 

USER nextjs

COPY --from=builder /app/public ./public

# Automatically leverage output traces to reduce image size
# https://nextjs.org/docs/advanced-features/output-file-tracing
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static


# Environment variables must be redefined at run time
COPY .env.production .env.production
ENV NEXT_TELEMETRY_DISABLED 1

CMD node server.js
