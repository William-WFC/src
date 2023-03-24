function findLong(str) {
    let obj = {
        oldArr: [],
        newArr: [],
    }
    let lastUpdateKey = 'oldArr' // 最后更新数组的key
    // 思路：整2个数组，把不重复的连续字符串赛进去，然后轮换比较，较长着保留，最后留下来的2个数组中的较长着就是最长不重复子串
    for (let i = 0; i < str.length; i++) {
        if (obj[lastUpdateKey].includes(str[i])) {
            const otherKey = lastUpdateKey === 'oldArr' ? 'newArr' : 'oldArr'
            if (obj[otherKey].length === 0) {
                obj[otherKey].push(str[i])
                lastUpdateKey = otherKey
            } else {
                if (obj[otherKey].length > obj[lastUpdateKey].length) {
                    obj[lastUpdateKey] = []
                    obj[lastUpdateKey].push(str[i])
                } else {
                    obj[otherKey] = []
                    obj[otherKey].push(str[i])
                    lastUpdateKey = otherKey
                }
            }
        } else {
            obj[lastUpdateKey].push(str[i])
        }
    }
    console.log(obj, lastUpdateKey)
    return (obj.oldArr.length > obj.newArr.length) ? obj.oldArr.length : obj.newArr.length
}

let str1 = 'abcabcbb', str2 = 'bbbb', str3 = 'pwwkew'
console.log(findLong('dvdf'))