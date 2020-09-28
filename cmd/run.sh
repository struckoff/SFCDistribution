POINTS=15000

./cmd -mode sfc -sfc.curve hilbert -sfc.curve.bits 2 -points $POINTS -clustermode 4 > geojsons/hilbert2x4_4c15k.json
./cmd -mode sfc -sfc.curve hilbert -sfc.curve.bits 3 -points $POINTS -clustermode 4 > geojsons/hilbert2x8_4c15k.json
./cmd -mode sfc -sfc.curve hilbert -sfc.curve.bits 4 -points $POINTS -clustermode 4 > geojsons/hilbert2x16_4c15k.json
./cmd -mode sfc -sfc.curve hilbert -sfc.curve.bits 5 -points $POINTS -clustermode 4 > geojsons/hilbert2x32_4c15k.json
./cmd -mode sfc -sfc.curve hilbert -sfc.curve.bits 6 -points $POINTS -clustermode 4 > geojsons/hilbert2x64_4c15k.json
./cmd -mode sfc -sfc.curve hilbert -sfc.curve.bits 7 -points $POINTS -clustermode 4 > geojsons/hilbert2x128_4c15k.json
./cmd -mode sfc -sfc.curve hilbert -sfc.curve.bits 8 -points $POINTS -clustermode 4 > geojsons/hilbert2x256_4c15k.json


./cmd -mode sfc -sfc.curve morton -sfc.curve.bits 2 -points $POINTS -clustermode 4 > geojsons/morton2x4_4c15k.json
./cmd -mode sfc -sfc.curve morton -sfc.curve.bits 3 -points $POINTS -clustermode 4 > geojsons/morton2x8_4c15k.json
./cmd -mode sfc -sfc.curve morton -sfc.curve.bits 4 -points $POINTS -clustermode 4 > geojsons/morton2x16_4c15k.json
./cmd -mode sfc -sfc.curve morton -sfc.curve.bits 5 -points $POINTS -clustermode 4 > geojsons/morton2x32_4c15k.json
./cmd -mode sfc -sfc.curve morton -sfc.curve.bits 6 -points $POINTS -clustermode 4 > geojsons/morton2x64_4c15k.json
./cmd -mode sfc -sfc.curve morton -sfc.curve.bits 7 -points $POINTS -clustermode 4 > geojsons/morton2x128_4c15k.json
./cmd -mode sfc -sfc.curve morton -sfc.curve.bits 8 -points $POINTS -clustermode 4 > geojsons/morton2x256_4c15k.json

./cmd -mode consistent -points $POINTS -clustermode 4 > geojsons/consistent_4c15k.json
