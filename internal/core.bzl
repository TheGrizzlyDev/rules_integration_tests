def _generate_java_client_impl(ctx):
    # From a template generates a simple to call java client that already wraps image and snapshot
    pass

generate_java_client = rule(
    _generate_java_client_impl,
)

def _generate_docker_criu_snapshot_impl(ctx):
    # Generates the actual snapshot of the container
    pass

generate_docker_criu_snapshot = rule(
    _generate_docker_criu_snapshot_impl,
)