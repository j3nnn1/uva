
 - ng new uva
 - deps
npm install @angular/animations@7.x  --save
npm install @angular/cdk@7.x   --save
npm install @angular/core@7.x   --save
npm install @angular/common@7.x  --save
npm install @angular/forms@7.x  --save
npm install @angular/platform-browser@7.x  --save
npm install @angular/platform-browser-dynamic@7.x  --save
npm install rxjs@6.x --save
npm install zone.js@~0.8.26 --save
npm install @angular/compiler@7.2.15  --save
 - npm install @swimlane/ngx-charts --save
 - npm install d3 --save npm install @types/d3-shape --save
 - npm install bootstrap -save
- deps
npm WARN @angular/flex-layout@8.0.0-beta.27 requires a peer of @angular/cdk@^8.0.0-rc.0 but none is installed. You must install peer dependencies yourself.
npm WARN @angular/flex-layout@8.0.0-beta.27 requires a peer of @angular/core@>=8.0.0-rc.5 but none is installed. You must install peer dependencies yourself.
npm WARN @angular/flex-layout@8.0.0-beta.27 requires a peer of @angular/common@>=8.0.0-rc.5 but none is installed. You must install peer dependencies yourself.
npm WARN @angular/flex-layout@8.0.0-beta.27 requires a peer of @angular/platform-browser@>=8.0.0-rc.5 but none is installed. You must install peer dependencies yourself.
npm WARN @swimlane/ng2-file-upload@1.0.1 requires a peer of @angular/common@^8.2.7 but none is installed. You must install peer dependencies yourself.
npm WARN @swimlane/ng2-file-upload@1.0.1 requires a peer of @angular/core@^8.2.7 but none is installed. You must install peer dependencies yourself.
npm WARN @swimlane/ngx-charts@12.0.1 requires a peer of @angular/cdk@7.x || 8.x but none is installed. You must install peer dependencies yourself.
npm WARN @swimlane/ngx-ui@27.1.0 requires a peer of @angular/cdk@^7.0.0 || ^8.0.0 but none is installed. You must install peer dependencies yourself.
npm WARN @swimlane/ngx-ui@27.1.0 requires a peer of codemirror@5.42.0 but none is installed. You must install peer dependencies yourself.
npm WARN @swimlane/ngx-ui@27.1.0 requires a peer of moment@^2.21.0 but none is installed. You must install peer dependencies yourself.
npm WARN @swimlane/ngx-ui@27.1.0 requires a peer of moment-timezone@^0.5.23 but none is installed. You must install peer dependencies yourself.
npm WARN @swimlane/ngx-ui@27.1.0 requires a peer of zone.js@^0.9.1 but none is installed. You must install peer dependencies yourself.
npm WARN bootstrap@4.3.1 requires a peer of jquery@1.9.1 - 3 but none is installed. You must install peer dependencies yourself.
npm WARN bootstrap@4.3.1 requires a peer of popper.js@^1.14.7 but none is installed. You must install peer dependencies yourself.

https://swimlane.gitbook.io/ngx-charts/installing

https://swimlane.github.io/ngx-charts/#/ngx-charts/line-chart

BUILD
=====
- ng build --prod --base-href "https://j3nnn1.github.io/uva/ui/dist/uva"

- https://j3nnn1.github.io/uva/ui/dist/uva/index.html

=====
git checkout -b gh-pages
🌹  git push origin gh-pages
🌹  npm install -g angular-cli-ghpages
🌹  ng build --prod --base-href https://[username].github.io/[repo]/
🌹  ngh --dir=dist/[project-name]
It is only necessary to set the the--base-href flag once, next time you build the project you can simply run:

ng build --prod
