<template>
  <div class="login-main">
    <div class="shell">
      <div class="box-left">
        <img src="@/assets/icon/star-eyes.png"
          style="width: 100px; height: 100px; padding-left: 0px;" />
        <span>欢迎使用电影售票系统</span>
      </div>
      <div class="box-right">
        <div>
          <div style="text-align: center; font-size: 200%; margin-top: 40px">
            欢迎登录
          </div>
          <el-form ref="form" :model="form" label-width="70px" style="margin-top: 30px">
            <el-form-item>
              <el-input prefix-icon="el-icon-user-solid" v-model="form.username" style="width: 250px"></el-input>
            </el-form-item>
            <el-form-item>
              <el-input prefix-icon="el-icon-lock" v-model="form.password" show-password
                style="width: 250px"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" round style="width: 100px" @click="login">
                登 录
              </el-button>
              <el-button type="danger" round style="width: 100px; margin-left: 50px" @click="register">
                注 册
              </el-button>
            </el-form-item>
          </el-form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import request from "../utils/request";
import sessionStorage from "../utils/storage";

export default {
  name: "Login",
  components: {},
  data() {
    return {
      form: {},
    };
  },
  methods: {
    login() {
      // request
      //   .post("http://localhost:9090/user/login", this.form)
      // 虚拟登陆菜单数据，需要请求接口时注释掉
      new Promise((resolve, rej) => {
        resolve({
          data: {
            role: "超级管理员"
          },
          code: 200
        })
      }).then((res) => {
        //暂时把路由权限关闭，需要关闭时注释
        sessionStorage.setItem('isLogin', true);
        if (res.data.status == 0) {
          this.$message({
            type: "error",
            message: "该账号已被封禁！如有疑问请联系管理员",
          });
          return false;
        } else if (res.code == 200) {//接口为200的时候跳转主页
          this.$message({
            type: "success",
            message: "登录成功！欢迎你，" + this.form.username + "！",
          });
          sessionStorage.setItem('isLogin', true)
          sessionStorage.setItem("user", JSON.stringify(res.data));
          this.$router.push("/home");
        } else {
          this.$message({
            type: "error",
            message: res.msg,
          });
        }
      }).catch(() => {
        sessionStorage.setItem('isLogin', true);
      });
    },
    register() {
      this.$router.push("/register");
    },
  },
};
</script>

<style scoped>
.login-main {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: url('../assets/icon/background.jpg') 100%;
}

* {
  padding: 0;
  margin: 0;
}

.shell {
  width: 640px;
  height: 320px;
  display: flex;
}

.box-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: rgba(255, 255, 255, 0.724);
  height: 280px;
  top: 20px;
  position: relative;
  width: 50%;
}

.box-left h2 {
  font: 900 50px "";
  margin: 50px 40px 40px;
}

.box-left span {
  /* margin: 40px; */
  margin-top: 20px;
  display: block;
  color: rgb(53, 171, 249);
  font-style: 14px;
  text-shadow: 0 0 3px #4875cf;
}

.box-right {
  background-color: #474a59;
  box-shadow: 0 0 40px 16px rgba(0, 0, 0, 0.2);
  color: #f1f1f2;
  width: 80%;
}
</style>
