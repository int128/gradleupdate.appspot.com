package service

import infrastructure.GitHubRepository

class GradleUpdateWorker {

    final templateRepository = new GitHubRepository('gradleupdate/GradleUpdateWorker')

    def bumpTemplate(String version) {
        def branch = "update-gradle-template-$version"
        templateRepository.removeBranch(branch)
        templateRepository.createBranch(branch, 'master')
    }

    def bumpUserRepository(String repo) {
        def branch = "update-gradle-of-$repo"
        templateRepository.removeBranch(branch)
        templateRepository.createBranch(branch, 'master')
    }

}