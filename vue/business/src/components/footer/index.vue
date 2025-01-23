<template>
  <a-layout-footer class="footer">
     <span>{{ Address }} <span v-if="Team">&amp;</span> <a :href="TeamSite" target="_blank">{{ Team }} </a> </span>
     <span class="copyright"> ⓒ 2018-{{nowyear}} <a :href="CompanySite" target="_blank">{{Company}}</a> {{ $t("footer.copyright") }}</span>
    </a-layout-footer>
</template>

<script lang="ts" setup>
 //获取网站配置-应用名称
 const Address = window?.globalConfig.Address
 const TeamSite = window?.globalConfig.TeamSite
 const Team = window?.globalConfig.Team
 const CompanySite = window?.globalConfig.CompanySite
 const Company = window?.globalConfig.Company
 const nowyear = new Date().getFullYear()
</script>

<style lang="less" scoped>
  .footer {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 40px;
    background-color: var(--color-fill-1);
    color: var(--color-text-2);
    text-align: center;
    a{
        color: #1890ff;
        text-decoration: none;
        background-color: transparent;
        outline: none;
        cursor: pointer;
        transition: color 0.3s;
    }
    .copyright{
      padding-left: 3px;
    }
  }
</style>
