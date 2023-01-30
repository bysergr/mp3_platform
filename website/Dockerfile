# Install dependencies only when needed
FROM node:16.18.0-alpine as deps
RUN apk add --no-cache libc6-compact
WORKDIR /app
COPY ["package.json", "package-lock.json*", "npm-shrinkwrap.json*", "./"]
RUN npm install --production --silent && mv node_modules ../

# Rebuild the source code only when needed
FROM node:16.18.0-alpine as builder 
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build 

# Production image, copy all the files and run next
FROM node:16.18.0-alpine as runner 
WORKDIR /app

ENV NODE_ENV production

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs 
EXPOSE 3000
ENV PORT 3000
CMD ["node", "server.js"]
