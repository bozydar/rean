### Installation

    $ git clone git@github.com:bozydar/rean.git
    $ cd rean
    $ make all
    $ make install

### Configuration

    $ cd your-project
    $ rean config --jira-username <username> --jira-password <jira-token> \
                --jira-url <jir-url> --jira-project-prefix <project-code>
E.g.
    
    $ rean config --jira-username bozydar@1centre.com --jira-password ****** \
                --jira-url https://1centre.atlassian.net --jira-project-prefix BT

It creates `.rean` file in the project directory

### Usage

**Show statuses** of all the jira tickets which are on `branch2` but not on `branch1`. Actually `git log branch1..branch2`. 


    $ rean show --from branch1 --to branch2

E.g.:

    $ rean show -f master -t deploy-staging
    7685: STAGING: Email methods to be removed (please check if these are still in prod)
    7101: STAGING: Sengrid template: ALN - APP escalated
    7530: READY TO TEST: 1CAF - CMM
    7659: RELEASED: 1Account issues
    5928: IN PROGRESS: Fraud prevention
    7008: IN PROGRESS: Email fixes - Batch 3
    7010: SELECTED FOR PROGRESSION: Sendgrid template: OWN - APP escalated
    7742: RELEASED: investigate - supplier unable to change ID
    7761: TESTING IN PROGRESS: 1CAD Bug x2
    7335: DONE: 1CAF - tidy up tasks
    7668: RELEASED: 1CAF Testing
    6824: STAGING: User profile - bulk reassign
    7011: SELECTED FOR PROGRESSION: Sendgrid template: OWN - APP deescalated
    

**Help**
    

    $ rean --help