import infrastructure.GitHub

import static util.RequestUtil.relativePath

final fromBranch = params.from_branch
final intoRepo = params.into_repo
final gradleVersion = params.gradle_version
assert fromBranch instanceof String
assert intoRepo instanceof String
assert gradleVersion instanceof String

final gitHub = new GitHub()

log.info("Creating a fork of $intoRepo")
final fork = gitHub.fork(intoRepo)

final fromUser = fork.owner.login
final fromRepo = fork.full_name
final intoBranch = fork.default_branch
assert fromUser instanceof String
assert fromRepo instanceof String
assert intoBranch instanceof String

log.info("Queue creating a branch $fromBranch on $fromRepo")
defaultQueue.add(
        url: relativePath(request, '4-branch.groovy'),
        params: [
                from_user: fromUser,
                from_branch: fromBranch,
                into_repo: intoRepo,
                into_branch: intoBranch,
                gradle_version: gradleVersion,
        ],
        countdownMillis: 1000)