# Launchpad module

this module governs all logic and activities related to launchpad's project. 

I take on a more wide understanding of launchpad. It should be a platform to release initial tokens to users.
1. ICO: ICO is both a way to release initial tokens and to raise funds.
1. lock drop: lock drop is a way to release initial tokens in exchange for user commitment.
1. air drop: air drop is just for releasing initial tokens to users.
1. liquidity bootstrapping auction: both as a way to release tokens, create initial liquidity, and gain user commitment.

I call these release mechanisms. Each release mechanism will have its own module.

A lot of effort goes to doing and researching the same release mechanisms again and again for each new project.

This module will save plenty of time for new project to lauch.

# a project structure
1. project_title: the title of release project
1. project_id: unique id of project
1. project_address: to store tokens
1. project_information: information for the project
1. start_time: start time of project schedule.
1. release_mechanism: the release mechanism used for this project. If a chain intends to use multiple release mechanism on different phases, it should create multiple projects.

# tx Message
1. CreateProject: tx message CreateProject to create a project.
1. DeleteProject: tx message DeleteProject to delete a project.
1. ModifyStartTime: tx message ModifyStartTime to modify start time of a project.
1. ModifyProjectInformation (through gov): gov proposal to modify a project information. A project information should be as stable as possible to build trust with users and investors.

# query Message
1. QueryProjectByID: query message project through project_id.