<script>
import {createApp} from 'vue';
const sfc = await import("./RepoContributors.vue")

export function initRepoContributorsChart() {
  const el = document.getElementById('repo-contributors-chart');
  if (el) {
    createApp(sfc).mount(el);
  }
}
</script>
