<template>
  <div class="cloud-sidebar">
    <div class="mian-bar">
      <!-- 展开和隐藏按钮 -->
      <div class="narrow-wrapper" @click="changeUnfoldStatus(1)" v-if="!isUnfold">
        <i class="iconfont icon-zhankaianniuicon narrow-icon"></i>
      </div>
      <div class="unfold-wrapper" @click="changeUnfoldStatus(0)" v-else>
        <i class="iconfont icon-icon01 unfold-icon"></i>
      </div>
      <!-- 展开和隐藏按钮 END -->

      <!-- 一级菜单列表 -->
      <ul class="first-menu-wrapper">
        <li class="item-menu-wrapper" v-for="(unfoldItem,unfoldItemIndex) in menuList" :key="unfoldItemIndex">
          <div class="item-menu-title" @click="unfoldItemMenu(unfoldItem,unfoldItemIndex)" >
            <img src="../assets/img/down.png" alt="" v-if="unfoldItemMenuIndex==unfoldItemIndex">
            <img src="../assets/img/right.png" alt="" v-else>
            <span v-if="!isUnfold">{{unfoldItem.mainTitle}}</span> 
          </div>
          <ul class="item-menu-list" :unfoldItemIndex="unfoldItemIndex" v-show="unfoldItemMenuIndex==unfoldItemIndex">
            <li v-for="(mainMenuItem,index) in unfoldItem.submenuList" :key="index" @click="computeSubMenuList(mainMenuItem)">
              <div class="submenu">
                <i class="iconfont" :class="mainMenuItem.icon"></i>
                <span v-if="!isUnfold">{{mainMenuItem.submenuTitle}}</span>
              </div>
            </li>
          </ul>
        </li>
      </ul>
      <!-- 一级菜单列表 END -->      
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isUnfold: false, // 判断展开还是隐藏一级菜单,
      unfoldItemMenuIndex: 0, // 判断展开还是隐藏一个菜单列表
      isShowMainMenu: false,
      menuArr: [
        {
          type: 1,
          menuList: [
            {
              main_memu_id: 0,
              mainTitle: "KEY",
              submenuList: [
                {
                  submenuTitle: "KV",
                  icon: "icon-yunguanjia",
                  submenuID: 1,
                  path: '/key/kv'
                },
                {
                  submenuTitle: "SRV",
                  icon: "icon-yuming",
                  submenuID: 2,
                  path: '/cloud/page2'
                }
              ]
            },
            {
              main_memu_id: 1,
              mainTitle: "大数据",
              submenuList: [
                {
                  submenuTitle: "机器学习",
                  submenuID: 3,
                  icon: "icon-xuexi",
                  path: '/cloud/page3'
                },
                {
                  submenuTitle: "数据集成",
                  submenuID: 4,
                  icon: "icon-shuju",
                  path: '/cloud/page5'
                }
              ]
            }
          ]
        }
      ],
      menuList: []
    };
  },
  methods: {
    // 顶级菜单折叠和展开
    changeUnfoldStatus(code) {
      if (code == 1) {
        this.isUnfold = true;
        sessionStorage.setItem("isUnfold", "1");
      } else {
        this.isUnfold = false;
        sessionStorage.setItem("isUnfold", "0");
      }
    },
    // 菜单选中
    unfoldItemMenu(unfoldItem, unfoldItemIndex) {
      if (this.unfoldItemMenuIndex == unfoldItemIndex) {
        this.unfoldItemMenuIndex = null;
      } else {
        this.unfoldItemMenuIndex = unfoldItemIndex;
      }
      sessionStorage.setItem("unfoldItemMenuIndex", unfoldItemIndex);
    },
    activateSubmenu() {
      let submenuObj = document.getElementsByClassName("submenu");
      for (let i = 0; i < submenuObj.length; i++) {
        submenuObj[i].onclick = function() {
          for (let j = 0; j < submenuObj.length; j++) {
            submenuObj[j].className = "submenu";
          }
          submenuObj[i].className = "submenu submenu-active";
          sessionStorage.setItem("mainMenuItemIndex", i);
        };
      }
    },
    computeSubMenuList(info) {
      // console.log(info)
      sessionStorage.setItem("submenuID", info.submenuID);
      this.subMenuList = info.submenuItemList;
      this.$router.push({ path: info.path });

    }
  },
  created() {
    // 通过不同权限显示不同页面
    let type = 1;
    this.menuArr.map(v => {
      if (v.type == type) {
        this.menuList = v.menuList;
      }
    });
  },
  mounted() {
    if (sessionStorage.getItem("isUnfold")) {
      if (sessionStorage.getItem("isUnfold") == 1) {
        this.isUnfold = true;
      } else {
        this.isUnfold = false;
      }
    }
    this.activateSubmenu();
    if (sessionStorage.getItem("mainMenuItemIndex")) {
      let index = Number(sessionStorage.getItem("mainMenuItemIndex"));
      let submenuObj = document.getElementsByClassName("submenu");
      submenuObj[index].className = "submenu submenu-active";
      this.unfoldItemMenuIndex = submenuObj[
        index
      ].parentNode.parentNode.getAttribute("unfoldItemIndex");
    }
    if (sessionStorage.getItem("submenuID")) {
      let submenuObj = document.getElementsByClassName("submenu");
      for (let i = 0; i < submenuObj.length; i++) {
        submenuObj[i].className = "submenu";
      }
      let index = Number(sessionStorage.getItem("submenuID"));
      submenuObj[index].className = "submenu submenu-active";
      this.menuList.map(v => {
        v.submenuList.map(vm => {
          if (vm.submenuID == Number(sessionStorage.getItem("submenuID")) + 1) {
            this.subMenuList = vm.submenuItemList;
          }
        });
      });
    }
  }

};
</script>

<style lang='scss'>
.cloud-sidebar {
  height: 100%;
  .mian-bar {
    float: left;
    height: 100%;
    background: #333744;
    .narrow-wrapper {
      width: 180px;
      height: 30px;
      background: #4a5064;
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      .narrow-icon {
        font-size: 12px;
        color: #aeafa7;
      }
    }
    .narrow-wrapper:hover .narrow-icon {
      color: #fff;
    }
    .unfold-wrapper {
      width: 50px;
      height: 30px;
      background: #4a5064;
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      .unfold-icon {
        font-size: 20px;
        color: #aeafa7;
      }
    }
    .unfold-wrapper:hover .unfold-icon {
      color: #fff;
    }
    .first-menu-wrapper {
      color: #fff;
      font-size: 12px;
      .item-menu-wrapper {
        .item-menu-title {
          display: flex;
          height: 40px;
          align-items: center;
          background: #42485b;
          cursor: pointer;
          img {
            margin-left: 18px;
            margin-right: 8px;
          }
        }
        .item-menu-title:hover {
          background: #00c1de;
        }
        .item-menu-list {
          .submenu {
            color: #aeb9c2;
            height: 40px;
            line-height: 40px;
            cursor: pointer;
            i {
              display: inline-block;
              width: 16px;
              height: 16px;
              margin: 0 17px;
            }
          }
          .submenu:hover {
            background: #4a5064;
          }
          .submenu-active {
            background: #00c1de;
            color: #fff;
          }
          .submenu-active:hover {
            background: #00c1de;
            color: #fff;
          }
        }
      }
    }
  }
}
</style>
