/**
 * 联动省市区插件
 * @author eecjimmy
 * @date 2015-12-13
 */
;(function ($) {
    $.fn.GangedArea = function (options) {
        self.opts = $.extend({}, $.fn.GangedArea.defaults, options);
        var $this = $(this);
        self.init = function () {
            $.getJSON(self.opts.url, function (data) {
                var html = '';
                if (self.opts.level > 0)html += self.getProvinceHtml(self.opts.provinceValue, data.province);
                if (self.opts.level > 1)html += self.getCityHtml(self.opts.provinceValue, self.opts.cityValue, data.city);
                if (self.opts.level > 2)html += self.getDistrict(self.opts.cityValue, self.opts.districtValue, data.district);
                $this.html(html);
                $this.on('change', 'select[name=' + self.opts.provinceName + ']', function () {
                    var val = $(this).val();
                    $(this).nextAll().remove();
                    var html = '';
                    if (self.opts.level > 1)html += self.getCityHtml(val, 0, data.city);
                    if (self.opts.level > 2)html += self.getDistrict(0, 0, data.district);
                    $(this).after(html);
                });
                $this.on('change', 'select[name=' + self.opts.cityName + ']', function () {
                    var val = $(this).val();
                    $(this).nextAll().remove();
                    var html = '';
                    if (self.opts.level > 2)html += self.getDistrict(val, 0, data.district);
                    $(this).after(html);
                })
            });
        };
        self.getProvinceHtml = function (provinceId, provinceData) {
            var html = '<select name="' + self.opts.provinceName + '" class="' + self.opts.className + '">';
            html += '<option value="">' + self.opts.tip + '</option>';
            var selectedHtml = '';
            $.each(provinceData, function (idx, item) {
                selectedHtml = provinceId == item.province_id ? ' selected' : '';
                html += '<option value="' + item.province_id + '"' + selectedHtml + '>' + item.province_name + '</option>';
            });
            html += '</select>';
            return html;
        };
        self.getCityHtml = function (provinceId, cityId, cityData) {
            var html = '<select name="' + self.opts.cityName + '" class="' + self.opts.className + '">';
            html += '<option value="">' + self.opts.tip + '</option>';
            var selectedHtml = '';
            $.each(cityData, function (idx, item) {
                if (item.province_id == provinceId) {
                    selectedHtml = cityId == item.city_id ? ' selected' : '';
                    html += '<option value="' + item.city_id + '"' + selectedHtml + '>' + item.city_name + '</option>';
                }
            });
            html += '</select>';
            return html;
        };
        self.getDistrict = function (cityId, districtId, districtData) {
            var html = '<select name="' + self.opts.districtName + '" class="' + self.opts.className + '">';
            html += '<option value="">' + self.opts.tip + '</option>';
            var selectedHtml = '';
            $.each(districtData, function (idx, item) {
                if (item.city_id == cityId) {
                    selectedHtml = districtId == item.district_id ? ' selected' : '';
                    html += '<option value="' + item.district_id + '"' + selectedHtml + '>' + item.district_name + '</option>';
                }
            });
            html += '</select>';
            return html;
        };
        self.init();
    };
    $.fn.GangedArea.defaults = {
        url: '', // json链接
        level: 3,// 深度
        className: '', // 附加样式
        provinceValue: '',// 默认省份id
        cityValue: '', // 默认市id
        districtValue: '', // 默认区id
        provinceName: 'province', // 省份name值
        cityName: 'city',//市name值
        districtName: 'district',// 区name值
        tip: '===请选择==='// 默认提示
    };
})(jQuery);