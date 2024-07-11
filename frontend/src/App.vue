<template>
  <!-- pake router view buat display view -->
  <router-view />
</template>

<script>
export default {
  name: 'App',
}
</script>

<style lang="scss">
/* Variable */
:root {
  --bg-primary: #fff;
  --bg-primary-brand: #f9d659;
  --bg-error: #ffc3c3;
  --bg-success: #c3ffdb;
  --text-primary: #333;
  --text-inverse: #fff;
  --text-error: #c81e1e;
  --text-success: #1ec84e;
  --border-primary: #dee2e6;
  --border-error: var(--text-error);
  --btn-primary: var(--bg-primary-brand);
  --btn-border-primary: var(--bg-primary-brand);
}

/* CSS Reset https://www.joshwcomeau.com/css/custom-css-reset/ */
*, *::before, *::after {
  box-sizing: border-box;
}
* {
  margin: 0;
}
html, body {
  height: 100%;
}
body {
  line-height: 1.5;
  -webkit-font-smoothing: antialiased;
}
img, picture, video, canvas, svg {
  display: block;
  max-width: 100%;
}
input, button, textarea, select {
  font: inherit;
}
p, h1, h2, h3, h4, h5, h6 {
  overflow-wrap: break-word;
}
#root, #__next {
  isolation: isolate;
}

/* base style */
body {
  font-family: sans-serif;
  color: var(--text-primary);
}
h1, h2, h3, h4 {
  line-height: 1.3;
}
h1 {
  font-size: 2rem;
}
h2 {
  font-size: 1.75rem;
}
h3 {
  font-size: 1.5rem;
}
h4 {
  font-size: 1.25rem;
}
#app {
  position: relative;
  height: 100%;
}
.btn {
  display: inline-block;
  padding: .75rem .75rem;
  color: var(--text-primary);
  text-align: center;
  text-decoration: none;
  vertical-align: middle;
  cursor: pointer;
  border: 0 solid transparent;
  border-radius: .5rem;
  background-color: transparent;
  transition: color .15s ease-in-out, background-color .15s ease-in-out, border-color .15s ease-in-out, box-shadow .15s ease-in-out;
  font-weight: 600;

  &-primary {
    background-color: var(--bg-primary-brand);
    border-color: var(--bg-primary-brand);
  }
}
.ornament {
  width: 100%;
  background: var(--bg-primary-brand);
  height: 40%;
  position: absolute;
  bottom: 0;
  left: 0;
  z-index: -1;
}
.container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  max-width: 1140px;
  width: 100%;
  margin: 0 auto;
  gap: 2rem;
  height: 100vh;
  padding: 0 2rem;
}

/* Custom Style */
.form {
  &-wrapper {
    margin-bottom: 2rem;

    label {
      margin-bottom: .25rem;
      display: block;
    }
  }

  &-control {
    display: block;
    width: 100%;
    padding: 1rem .75rem;
    font-size: 1rem;
    font-weight: 400;
    line-height: 1.5;
    color: var(--text-primary);
    appearance: none;
    background-color: var(--bg-primary);
    background-clip: padding-box;
    border: 1px solid var(--border-primary);
    border-radius: .5rem;
    transition: all .15s ease-in-out;

    &:focus, &:focus-visible {
      outline: none;
      border-color: var(--bg-primary-brand);
    }
  }
}

.left {
  .title {
    font-size: 3rem;
  }
  .subtitle {
    font-size: 1.5rem;
    font-weight: 400;
  }
}

.right {
  padding: 2rem;
  border: 1px solid #e7e7e7;
  border-radius: 1rem;
  box-shadow: 0 0 16px rgb(0 0 0 /10%);
  background: var(--bg-primary);
  max-width: 460px;
  width: 100%;

  button {
    margin-bottom: 1rem;
  }

  .other-way {
    text-align: center;
  }
}

.card-title {
  margin-bottom: 1.5rem;
  text-align: center;
  text-transform: uppercase;
}

.input-error {
  border-color: var(--tex);
}

.helper-text {
  font-size: .875rem;

  &.error {
    color: var(--text-error);
  }
}

.notification {
  padding: .5rem .75rem;
  border-radius: .25rem;
  margin-top: 1rem;
  display: flex;
  gap: .25rem;
  align-items: center;

  .icon {
    width: 28px;
    height: 28px;

    &-info {
      mask: url('./assets/icn-info.svg') no-repeat center/contain;
    }
  }

  &.success {
    background: var(--bg-success);
    color: var(--text-success);

    .icon-info {
      background-color: var(--text-success);
    }
  }

  &.error {
    background: var(--bg-error);
    color: var(--text-error);

    .icon-info {
      background-color: var(--text-error);
    }
  }
}

@media screen and (max-width: 768px) {
  .container {
    flex-direction: column;
    height: 100%;
    padding-block: 1.5rem;
  }
  .left {
    text-align: center;
  }
}
</style>
