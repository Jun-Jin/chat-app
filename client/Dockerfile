# build stage
FROM node:20.10.0-alpine3.18 as react-build
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
RUN ls -l /usr/src/app

# runtime stage
FROM nginx:1.21.3-alpine
COPY --from=react-build /usr/src/app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

# docker build -t chat-app-client:0.0.1 .
# docker run -d -p 3000:80 --name chat-app-client chat-app-client:0.0.1
