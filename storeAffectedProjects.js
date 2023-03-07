const fs = require('fs');

function storeAffectedProjects() {
  const affectedProjects = JSON.parse(fs.readFileSync('./affected-projects.json'));
  const projectNames = affectedProjects.projects.map(p => p.name);
  process.env.PROJECTS = projectNames.join(',');
}
storeAffectedProjects();
