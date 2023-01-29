<template>
  <div class="rounded-lg font-bold text-black px-10 pb-16 pt-4 underline decoration-2 decoration-pink-500 mb-8" id="main-cta">
    <p class="drop-shadow drop-shadow-2xl text-3xl">Get out and about</p>
    <div class="drop-shadow drop-shadow-2xl flex">
      <p class="text-sm mt-1 mr-1 tracking-wide">Book a flight with TimelyTravel</p>
      <i class="fa-solid fa-arrow-right mt-1.5"></i>
    </div>
  </div>
</template>

<script>
export default {
  name: "CTA"
};
</script>

<style scoped>
#main-cta {
  background-image: url("@/assets/home.jpg");
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
}
</style>