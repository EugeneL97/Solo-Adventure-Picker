# Build stage
FROM node:20-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy rest of the application
COPY . .

# Expose port 5173 (Vite's default port)
EXPOSE 5173

# Start the application
CMD ["npm", "run", "dev", "--", "--host"]