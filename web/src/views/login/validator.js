export default {
    useradmin: [
        {
            required: true,
            message: '请填写账号',
            trigger: 'blur',
        }  
    ],
    password: [
        {
            required: true,
            message: '请填写密码',
            trigger: 'blur',
        },
        {
            min: 4,
            max: 12,
            message: '密码至少大于4位 不超过 12位',
            trigger: 'blur',
        },
    ],

    captcha: [
        {
            required: true,
            message: '请填写验证码',
            trigger: 'blur',
        },
        {
            min: 4,
            message: '验证码至少四位',
            trigger: 'blur',
        },
    ]
}