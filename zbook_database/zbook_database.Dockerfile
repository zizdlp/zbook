# Build stage
FROM postgres:16.2-alpine3.19 as builder
WORKDIR /app

COPY deps_postgres.sh .
RUN chmod +x deps_postgres.sh  \
  && sh deps_postgres.sh 
RUN git clone --recurse-submodules https://github.com/jaiminpan/pg_jieba
WORKDIR /app/pg_jieba
RUN mkdir build && \ 
  cd build && \
  cmake .. && \
  make && \
  make install 
WORKDIR /app

# Run stage
FROM postgres:16.2-alpine3.19

# Copy built extension from builder stage
COPY --from=builder /usr/local /usr/local

# Set the shared_preload_libraries configuration
CMD ["postgres", "-c", "shared_preload_libraries=/usr/local/lib/postgresql/pg_jieba.so"]