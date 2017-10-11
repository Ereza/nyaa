{{ extends "layouts/index_site" }}
{{block title()}}{{ T("upload")}}{{end}}
{{block content_body()}}
<div style="text-align: left;" class="box">
	<h1 style="text-align:center;margin-top; 4px;">Upload status</h1>
	<table class="multiple-upload">
		<tbody>
			<tr><td colspan="4"><h3>Nyaa Pantsu upload status</h3></td></tr>
			<tr class="upload-status"><td class="upload-site-name">Nyaa Pantsu</td><td>Status:</td><td class="icon-finished finished"></td><td class="upload-progress finished">Finished</td></tr>
			<tr><td colspan="4" class="uploaded-url"><input type="text" class="form-input" value="{{if Sukebei()}}{{Config.WebAddress.Sukebei}}{{else}}{{Config.WebAddress.Nyaa}}{{end}}/view/{{UploadMultiple.PantsuID}}" disabled></td></tr>
			
			{{ if UploadMultiple.NyaasiStatus != 0 }}
			<tr><td colspan="4"><h3>Nyaa.si upload status</h3></td></tr>
			<tr class="upload-status"><td class="upload-site-name">Nyaa.si</td><td>Status:</td><td class="icon-{{ if UploadMultiple.NyaasiStatus == 1}}pending pending{{else if UploadMultiple.NyaasiStatus == 2}}error error{{else}}finished finished{{end}}"></td><td class="upload-progress {{ if UploadMultiple.NyaasiStatus == 1}}pending{{else if UploadMultiple.NyaasiStatus == 2}}error{{else}}finished{{end}}">{{ if UploadMultiple.NyaasiStatus == 1}}Pending{{else if UploadMultiple.NyaasiStatus == 2}}Error{{else}}Finished{{end}}</td></tr>
			<tr><td colspan="4" class="uploaded-url"><input type="text" class="form-input" value="{{UploadMultiple.NyaasiMessage}}" disabled></td></tr>
			{{end}}
			
			{{ if UploadMultiple.AnidexStatus != 0 }}
			<tr><td colspan="4"><h3>Anidex upload status</h3></td></tr>
			<tr class="upload-status"><td class="upload-site-name">Anidex</td><td>Status:</td><td class="icon-{{ if UploadMultiple.AnidexStatus == 1}}pending pending{{else if UploadMultiple.AnidexStatus == 2}}error error{{else}}finished finished{{end}}"></td><td class="upload-progress {{ if UploadMultiple.AnidexStatus == 1}}pending{{else if UploadMultiple.AnidexStatus == 2}}error{{else}}finished{{end}}">{{ if UploadMultiple.AnidexStatus == 1}}Pending{{else if UploadMultiple.AnidexStatus == 2}}Error{{else}}Finished{{end}}</td></tr>
			<tr><td colspan="4" class="uploaded-url"><input type="text" class="form-input" value="{{UploadMultiple.AnidexMessage}}" disabled></td></tr>
			{{end}}
			
			{{ if UploadMultiple.TToshoStatus != 0 }}
			<tr><td colspan="4"><h3>Tokyo Tosho upload status</h3></td></tr>
			<tr class="upload-status"><td class="upload-site-name">Anidex</td><td>Status:</td><td class="icon-{{ if UploadMultiple.TToshoStatus == 1}}pending pending{{else if UploadMultiple.TToshoStatus == 2}}error error{{else}}finished finished{{end}}"></td><td class="upload-progress {{ if UploadMultiple.TToshoStatus == 1}}pending{{else if UploadMultiple.TToshoStatus == 2}}error{{else}}finished{{end}}">{{ if UploadMultiple.TToshoStatus == 1}}Pending{{else if UploadMultiple.TToshoStatus == 2}}Error{{else}}Finished{{end}}</td></tr>
			<tr><td colspan="4" class="uploaded-url"><input type="text" class="form-input" value="{{UploadMultiple.TToshoMessage}}" disabled></td></tr>
			{{end}}
		</tbody>
	</table>
 
</div>
{{end}}
{{ block footer_js()}}
<script type="text/javascript" src="/js/translation.js?v={{ Config.Version}}{{ Config.Build }}"></script>
<script type="text/javascript" src="/js/upload.js?v={{ Config.Version}}{{ Config.Build }}"></script>
{{end}}
