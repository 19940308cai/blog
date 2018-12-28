/**
 * 地区选择类
 */
(function($) {
    $.Area = {
        settings: {
            insertId: "area",
            nameList: ['provincename', 'cityname', 'areaname'],
            idList: ['provinceid', 'cityid', 'areaid'],
            isdisplay: false,
            defaultArea: [],
            limit: [],
            getAreaUrl: '/',
            func: function(obj) {}
        },

        init: function(settings) {
            if (settings) {
                $.extend(this.settings, settings);
            }
            this.initChinaAreaSelect();
            return this;
        },

        initChinaAreaSelect: function() {
            var settingsObj = this.settings;
            var thisObj = this;
            var selectHtml = '<select id="' + settingsObj.idList[0] + '" name="' + settingsObj.nameList[0] + '" style="display:none;" class="selHeight"></select>' + "\n";
            selectHtml += '<select id="' + settingsObj.idList[1] + '" name="' + settingsObj.nameList[1] + '" style="display:none;" class="selHeight"></select>' + "\n";
            selectHtml += '<select id="' + settingsObj.idList[2] + '" name="' + settingsObj.nameList[2] + '" style="display:none;" class="selHeight"></select>' + "\n";

            $("#" + settingsObj.insertId).html(selectHtml);
            this.getAreaSelectHtml('0', settingsObj, 0);

            // 绑定事件
            $("#" + settingsObj.idList[0]).change(function() {
                if (0 == $("#" + settingsObj.idList[0]).val()) {
                    $("#" + settingsObj.idList[1]).val(0).hide();
                    $("#" + settingsObj.idList[2]).val(0).hide();
                    return false;
                }
                $("#" + settingsObj.idList[1]).val(0);
                $("#" + settingsObj.idList[2]).val(0).hide();
                thisObj.getAreaSelectHtml($(this).val(), settingsObj, 1);
                settingsObj.func(this);
            });

            $("#" + settingsObj.idList[1]).change(function() {
                if (0 == $("#" + settingsObj.idList[1]).val()) {
                    $("#" + settingsObj.idList[2]).val(0).hide();
                    return false;
                }
                $("#" + settingsObj.idList[2]).val(0);
                thisObj.getAreaSelectHtml($(this).val(), settingsObj, 2);
                settingsObj.func(this);
            });

            $("#" + settingsObj.idList[2]).change(function() {
                settingsObj.func(this);
            });

            if (settingsObj.defaultArea.length > 0) {
                for (var i = 0; i < 3; i++) {
                    if (0 != i) {
                        thisObj.getAreaSelectHtml(settingsObj.defaultArea[i - 1], settingsObj, i);
                    }
                    $("#" + settingsObj.idList[i]).val(settingsObj.defaultArea[i]);
                }
            }
        },

        getAreaSelectHtml: function(parentid, settingsObj, key) {
            selectName = settingsObj.nameList[key];
            selectId = settingsObj.idList[key];
            limits = settingsObj.limit[key];

            if (!selectName) {
                return false;
            }

            parentid = parentid || '0';
            selectId = selectId || 'province';

            if (settingsObj.isdisplay) {
                // 隐藏二级地区
                var provinceToCity = {
                    '1': [3302, '北京市'],
                    '2': [3303, '天津市'],
                    '9': [3304, '上海市'],
                    '22': [3305, '重庆市']
                };
                if ('undefined' != typeof(provinceToCity[parentid])) {
                    $("#" + selectId).html('<option class="optHeight" class="optHeight"   value="' + provinceToCity[parentid][0] + '">' + provinceToCity[parentid][1] + '</option>' + "\n");
                    $("#" + selectId).hide();
                    parentid = provinceToCity[parentid][0];
                    selectName = settingsObj.nameList[key + 1];
                    selectId = settingsObj.idList[key + 1];
                    limits = settingsObj.limit[key + 1];
                }
            }

            var selectHtml = '';

            $.ajax({
                type: "GET",
                async: false,
                url: settingsObj.getAreaUrl,
                data: "parentid=" + parentid,
                dataType: "JSON",
                success: function(arrList) {
                    if (null == arrList) {
                    	alert("数据错误");
                        return;
                    }

                    selectHtml += '<option class="optHeight" value="0">--请选择--</option>' + "\n";

                    if ('undefined' != typeof(limits)) {
                        for (var v in arrList) {
                            if (-1 != $.inArray(parseInt(v), limits)) {
                                selectHtml += '<option class="optHeight" value="' + v + '">' + arrList[v] + '</option>' + "\n";
                            }
                        }
                    } else {
                        for (var v in arrList) {
                            selectHtml += '<option class="optHeight" value="' + v + '">' + arrList[v] + '</option>' + "\n";
                        }
                    }

                    $("#" + selectId).html(selectHtml);
                    $("#" + selectId).show();
                }
            });
        },
        setEmpty: function(selectId) {
            selectId = selectId || 'province';
            $("#" + selectId).html('<option class="optHeight" value="0">--请选择--</option>' + "\n");
            $("#" + selectId).hide();
        }
    }
})(jQuery);
