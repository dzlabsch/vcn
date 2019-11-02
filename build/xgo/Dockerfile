FROM techknowlogick/xgo:go-1.13.x

# Inject the customized build script
ADD build.sh /build.sh
ENV BUILD /build.sh
RUN chmod +x $BUILD

ENTRYPOINT ["/build.sh"]
