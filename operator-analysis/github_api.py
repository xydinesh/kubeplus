#!/usr/bin/env python3

from github import Github
import os
from logzero import logger


def collect_operators_runner():
    """ Runner entry point, outputs alloperators.txt file."""
    logger.info(f'Using github api to collect all operators')
    logger.info('For api access, reading form $USER and $PASSWORD env vars..')
    logger.info('Please make sure these are set.')
    user = os.getenv("USER")
    passwd = os.getenv("PASSWORD")
    g = Github(user, passwd)
    pagin_of_repo = g.search_repositories("kubernetes+operators")
    with open("alloperators.txt", 'w') as outf:
        for repo in pagin_of_repo:
            # name = repo.name
            clone_url = repo.clone_url
            last_commit = repo.pushed_at
            num_contributors = repo.get_contributors().totalCount
            stars = repo.stargazers_count

            outf.write("clone_url:%s, last_commit:%s, \
num_contributors:%s, \
stars:%s\n" % (clone_url,
                last_commit, num_contributors,
                stars))
