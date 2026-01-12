## DockerFile: For "Nginx" Application ##

# Start from the latest "nginx" base image (alpine)
FROM nginx:alpine

# Remove the default "Nginx" configuration 'file'
RUN rm /usr/share/nginx/html/index.html

# Copy the custom 'index.html' to the "Nginx" web 'root'
COPY index.html /usr/share/nginx/html/

# Expose the port "80" to 'access' the web 'page'
EXPOSE 80

# Start the "Nginx" server
CMD ["nginx", "-g", "daemon off;"]
