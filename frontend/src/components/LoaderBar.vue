<template>
  <div class="w-full flex justify-center items-center my-4">
    <div class="loader-bar h-[7px] w-2/3 rounded-full overflow-hidden relative bg-gradient-to-r from-[#232526] to-[#414345]">
      <div class="loader-gradient absolute left-0 top-0 h-full w-1/3 animate-loader-gradient" />
    </div>
  </div>
</template>

<style scoped>
@keyframes loader-gradient {
  0% {
    left: 0%;
    right: auto;
    background: linear-gradient(270deg, #ffe600 0%, #232526 100%);
  }
  50% {
    left: 67%;
    right: auto;
    background: linear-gradient(90deg, #232526 0%, #ffe600 100%);
  }
  100% {
    left: 0%;
    right: auto;
    background: linear-gradient(270deg, #ffe600 0%, #232526 100%);
  }
}
.loader-gradient {
  background: linear-gradient(270deg, #ffe600 0%, #232526 100%);
  width: 33%;
  height: 100%;
  border-radius: 9999px;
  animation: loader-gradient 1.5s cubic-bezier(0.4,0.0,0.2,1) infinite;
}
</style>
