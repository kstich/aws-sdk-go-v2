package software.amazon.smithy.aws.go.codegen;

import java.util.List;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.go.codegen.GoDelegator;
import software.amazon.smithy.go.codegen.GoSettings;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SymbolUtils;
import software.amazon.smithy.go.codegen.integration.GoIntegration;
import software.amazon.smithy.go.codegen.integration.MiddlewareRegistrar;
import software.amazon.smithy.go.codegen.integration.RuntimeClientPlugin;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.ServiceShape;
import software.amazon.smithy.utils.ListUtils;

/**
 * wraps a smithy request error with an AWS Request error.
  */
public class AWSResponseErrorWrapper implements GoIntegration {
    private static final String ADD_ERROR_WRAPPER = "addResponseErrorWrapper";
    private static final String ADD_ERROR_WRAPPER_INTERNAL = "AddResponseErrorWrapper";

    /**
     * Gets the sort order of the customization from -128 to 127, with lowest
     * executed first.
     *
     * @return Returns the sort order, defaults to -40.
     */
    @Override
    public byte getOrder() {
        return 127;
    }

    @Override
    public void writeAdditionalFiles(
            GoSettings settings,
            Model model,
            SymbolProvider symbolProvider,
            GoDelegator goDelegator
    ) {
        ServiceShape service = settings.getService(model);

        goDelegator.useShapeWriter(service, this::writeMiddlewareHelper);
    }

    private void writeMiddlewareHelper(GoWriter writer) {
        writer.openBlock("func $L(stack *middleware.Stack) {", "}", ADD_ERROR_WRAPPER, () -> {
            writer.write("$T(stack)",
                    SymbolUtils.createValueSymbolBuilder(ADD_ERROR_WRAPPER_INTERNAL,
                            AwsGoDependency.AWS_HTTP_TRANSPORT).build()
            );
        });
        writer.insertTrailingNewline();
    }

    @Override
    public List<RuntimeClientPlugin> getClientPlugins() {
        return ListUtils.of(
                RuntimeClientPlugin.builder()
                        .registerMiddleware(MiddlewareRegistrar.builder()
                                .resolvedFunction(SymbolUtils.createValueSymbolBuilder(ADD_ERROR_WRAPPER).build())
                                .build())
                        .build()
        );
    }
}
