package zjwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	minSize = 0
}

func TestCompactor(t *testing.T) {
	t.Run("Small", func(t *testing.T) {
		jwt := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG7lIiwiaWF0IjoxNTE2MjM5MDIyfQ.EpM5XBzTJZ4J8AfoJEcJrjth8pfH28LWdjLo90sYb9g`
		compactJWT, err := ZJWT(jwt)
		assert.NoError(t, err)
		expandedJWT, err := JWT(compactJWT)
		assert.NoError(t, err)
		assert.Equal(t, jwt, expandedJWT)
		assert.Equal(t, jwt, compactJWT)
		assert.Equal(t, 100, 100*len(compactJWT)/len(jwt))
	})
	t.Run("Large", func(t *testing.T) {
		jwt := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJncm91cHMiOlsiU0JHIiwiU0JHOk1lbWJlcnMiLCJTQkc6dWkiLCJTQkc6emlvbiIsIlNCRzpoYXJtb255LXVpLWNvbXBvbmVudHMtZXh0ZW5kZWQiLCJTQkc6bmV4Z2VuLXNlcnZpY2VzLXRlYW0iLCJTQkc6c2JnLXFhLXdyaXRlIiwiU0JHOmFkbWluLXFiYS11aS1hdXRvbWF0aW9uIiwiU0JHOnNiZy9xYm8tYXV0b21hdGlvbi10ZWFtLWFkbWluIiwiU0JHOmdpdGxhYi10b29scy13cml0ZSIsIlNCRzpTQkctT2ZmZXJpbmdzLVFFIiwiU0JHOnFiby1iaWxsaW5nLXRlc3QtcWUtdGVhbSIsIlNCRzpxYm8tYmlsbGluZy10ZXN0LWFkbWluIiwiU0JHOnFiYS1xdWFsaXR5LXRlYW0iLCJTQkc6cWJvLWZ0dS1xdWFsaXR5LWFkbWluLXRlYW0iLCJTQkc6cWJhLW1pZ3JhdGlvbi10b29scy10ZWFtIiwiU0JHOlhjZWxsZW5jZUZvcnVtSlMiLCJTQkc6dHJpbml0eWpzLW1vYmlsZSIsIm1vbmV5LWFuZC1tYWdpYyIsIm1vbmV5LWFuZC1tYWdpYzpPd25lcnMiLCJtb25leS1hbmQtbWFnaWM6TSAmIE0gU2NydW0gVGVhbSIsInBheXJvbGwiLCJwYXlyb2xsOkRFUFJFQ0FURURfT3duZXJzIiwicGF5cm9sbDpQYXlyb2xsQWxsIiwiU0JHLVBDUy1DSUNEIiwiU0JHLVBDUy1DSUNEOk93bmVycyIsIlNCLVFCTy1RRS1PZmZlcmluZ3MiLCJTQi1RQk8tUUUtT2ZmZXJpbmdzOk93bmVycyIsIlNCLVFCTy1RRS1PZmZlcmluZ3M6cWJvLWxxYS1hZG1pbiIsIlNCLVFCTy1RRS1PZmZlcmluZ3M6UUUtR3JvdXAiLCJTQi1RQk8tUUUtUGF5bWVudHMiLCJTQi1RQk8tUUUtUGF5bWVudHM6T3duZXJzIiwiU0ItUUJPLVFFLVBheW1lbnRzOnFiby1wYXltZW50cy1xZS10ZWFtIiwiU0ItUUJPLVFFLVBheXJvbGwiLCJTQi1RQk8tUUUtUGF5cm9sbDpPd25lcnMiLCJTQi1RQk8tUUUtUGF5cm9sbDpTQi1RQk8tUUUtUGF5cm9sbCIsIlNCLVFCTy1RRS1QYXlyb2xsOnBheXJvbGwtcWUtdGVhbSIsIml0YWciLCJpdGFnOml0YWctdWkiLCJpdGFnOml0YWctd3MiLCJpdGFnOm1hcnRpbmktanMiLCJnaXRsYWItbWlncmF0aW9uLTAwMDEiLCJnaXRsYWItbWlncmF0aW9uLTAwMDE6T3duZXJzIiwiZ2l0bGFiLW1pZ3JhdGlvbi0wMDAyIiwiZ2l0bGFiLW1pZ3JhdGlvbi0wMDAyOk93bmVycyIsImdpdGxhYi1taWdyYXRpb24tMDA0NSIsImdpdGxhYi1taWdyYXRpb24tMDA0NTpPd25lcnMiLCJnaXRsYWItbWlncmF0aW9uLTAwNDYiLCJnaXRsYWItbWlncmF0aW9uLTAwNDY6T3duZXJzIiwiU0JHLVBsdWdpbnMiLCJTQkctUGx1Z2luczpDb3JlIiwiU0JHLVBsdWdpbnM6aW1wcm92ZWQtaW52ZW50b3J5IiwiU0JHLVBsdWdpbnM6d2ludm9pY2UtZ21haWwiLCJBbHRpbWV0cmlrIiwiQWx0aW1ldHJpazpBbHRpbWV0cmlrIiwiU0ItUUJPLVFFLU9mZmVyaW5ncy1Nb2JpbGUiLCJTQi1RQk8tUUUtT2ZmZXJpbmdzLU1vYmlsZTpPd25lcnMiLCJRQk8tQ29yZSIsIlFCTy1Db3JlOnFiby1qYXZhLW1vbm9saXRoLXdyaXRlLWFjY2VzcyIsIlFCTy1JbnZlbnRvcnkiLCJRFk8tSW52ZW50b3J5OkFkbWlucyIsIlFCTy1JbnZlbnRvcnk6Q29udHJpYnV0b3JzIiwiUUJPLUludmVudG9yeTpDSUNEIiwiUUJPLUJpbGxpbmciLCJRQk8tQmlsbGluZzpCaWxsaW5nLVVJLUFkbWluIiwiU0JHLVFFIiwiZmFicmljLWFwcC1zaGVsbHMiLCJmYWJyaWMtYXBwLXNoZWxsczpNb2JpbGUgU2hlbGwiLCJmYWJyaWMtYXBwLXVpLWNvbXBvbmVudHMiLCJmYWJyaWMtYXBwLXVpLWNvbXBvbmVudHM6ZmFicmljLWFwcC11aS1jb21wb25lbnRzLW1lbWJlcnMiLCJmYWJyaWMtYXBwLXVpLWNvbXBvbmVudHM6SURTIEV4dGVuZGVkIENvcmUgVGVhbSIsImRldi10ZXN0IiwiZGV2LXRlc3Q6VEVQIiwidGVzdC1wbGF0Zm9ybSIsInRlc3QtcGxhdGZvcm06dGVzdC1wbGF0Zm9ybS1vcmctYWRtaW5zIiwidGVzdC1wbGF0Zm9ybTpvdmVyd2F0Y2gtY29yZSIsIm9wZW4tdWktY29tcG9uZW50cyIsIlNCU0VHLUVQSUMiLCJhY2NvdW50aW5nLWNvcmUiLCJhY2NvdW50aW5nLWNvcmU6YWNjb3VudGluZy1jb3JlLXFiby1jYXNlY2VudGVyLXVpIiwia3ViZXJuZXRlcyIsImt1YmVybmV0ZXM6c2Jnc2VnLWNkcC10ZWFtIiwic3ZjLXNic2VnLWNkcCIsIlNVRFMiLCJTVURTOlNVRFMiLCJhcHAtc2hlbGwiLCJhcHAtc2hlbGw6YXJnb2NkLWFkbWlucyIsIlNCRy1UcmlhZ2VCb3QiLCJzYW5kYm94LXNhbmRib3giLCJzYW5kYm94LXNhbmRib3g6dXgtd29ya3Nob3AtcmFqIiwic2FuZGJveC1zYW5kYm94OnV4LXdzLXJhaiIsInNhbmRib3gtc2FuZGJveDpyYWotdGVzdC13cyIsInBheXJvbGwtcGF5Y2hlY2siLCJwYXlyb2xsLXBheWNoZWNrOlNCR1NDVEVQIiwiYXBwLXVpY29tcG9uZW50cyIsImFwcC11aWNvbXBvbmVudHM6UGxheWdyb3VuZC1ydmFzaWthcmxhIiwiYXBwLXVpY29tcG9uZW50czphcHAtdWljb21wb25lbnRzLWlkcy1kYXNoYm9hcmQtdWkiLCJhcHAtdWljb21wb25lbnRzOmlkcy1wbGF5Z3JvdW5kLXVpIiwiVVgtSW5mcmEiLCJVWC1JbmZyYTpDb3JlIiwiZGVzaWduLXN5c3RlbXMiXX0.XDozAEiz9AIVkA3DFbPOKG6msM43gT5zup3oxsHg_4Q`
		compactJWT, err := ZJWT(jwt)
		assert.NoError(t, err)
		expandedJWT, err := JWT(compactJWT)
		assert.NoError(t, err)
		assert.Equal(t, jwt, expandedJWT)
		assert.Equal(t, 37, 100*len(compactJWT)/len(jwt))
	})
}

func TestCompact(t *testing.T) {
	_, err := ZJWT(".")
	assert.Error(t, err)
}

func TestExpand(t *testing.T) {
	_, err := JWT(".")
	assert.Error(t, err)
	_, err = JWT("...")
	assert.Error(t, err)
	_, err = JWT("zJWT/v1:...")
	assert.Error(t, err)
	_, err = JWT("zJWT/v1:..!.")
	assert.Error(t, err)
}
