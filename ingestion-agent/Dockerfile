FROM php:7.1-apache
COPY ./src/ /var/www/html/



RUN pecl install redis && docker-php-ext-enable redis







