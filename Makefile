OUT_DIR = out

GO ?= go
JAVAC ?= javac
JAVA ?= java
RM ?= rm

.PHONY: run
run:
	@$(RM) -rf $(OUT_DIR)
	@$(JAVAC) -h $(OUT_DIR) -d $(OUT_DIR) java/Main.java
	@CGO_CFLAGS="-I ${JAVA_HOME}/include -I ${JAVA_HOME}/include/darwin" $(GO) build -buildmode=c-shared -o $(OUT_DIR)/libgreetings.dylib ./internal/jni_export/main.go
	@$(JAVA) -Djava.library.path=$(OUT_DIR) -cp $(OUT_DIR) Main

clean:
	@$(RM) -rf $(OUT_DIR)