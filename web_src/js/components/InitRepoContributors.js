import {createApp} from 'vue';

export async function initRepoContributorsChart() {
  const el = document.getElementById('repo-contributors-chart');
  if (!el) {
    return
  }
  const sfc = await import("./RepoContributors.vue")
  createApp(sfc).mount(el);
}
