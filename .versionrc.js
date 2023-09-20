module.exports = {
  // hack to skip recursive build loop in github actions
  releaseCommitMessageFormat: 'chore(release): {{currentTag}} [skip ci]',
}
