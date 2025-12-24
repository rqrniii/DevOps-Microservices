#!/bin/bash

# 🎨 CSS Build Verification Script
# Run this to verify Tailwind CSS is building correctly

set -e

echo "🎨 Tailwind CSS v4 Build Verification"
echo "======================================"
echo ""

# Navigate to frontend directory
cd frontend/todo-frontend/

echo "Step 1: Clean previous build..."
rm -rf dist
echo "✓ Cleaned"
echo ""

echo "Step 2: Building with Bun..."
bun run build
echo "✓ Build complete"
echo ""

echo "Step 3: Checking for CSS files..."
CSS_FILES=$(find dist -name "*.css" -type f)

if [ -z "$CSS_FILES" ]; then
    echo "❌ ERROR: No CSS files found in dist!"
    echo ""
    echo "This might indicate a problem with Tailwind CSS processing."
    exit 1
fi

echo "✓ CSS files found:"
echo "$CSS_FILES"
echo ""

echo "Step 4: Checking CSS file size..."
for css_file in $CSS_FILES; do
    SIZE=$(wc -c < "$css_file")
    echo "  - $(basename $css_file): ${SIZE} bytes"
    
    if [ $SIZE -lt 1000 ]; then
        echo "    ⚠️  WARNING: CSS file is very small (< 1KB)"
        echo "    This might mean Tailwind classes aren't being processed"
    else
        echo "    ✓ Size looks good"
    fi
done
echo ""

echo "Step 5: Checking if Tailwind classes are in CSS..."
SAMPLE_CSS=$(cat $CSS_FILES | head -100)

# Check for common Tailwind utilities
if echo "$SAMPLE_CSS" | grep -q "bg-gradient-to"; then
    echo "✓ Found Tailwind gradient classes"
elif echo "$SAMPLE_CSS" | grep -q "flex"; then
    echo "✓ Found Tailwind utility classes"
else
    echo "⚠️  WARNING: Couldn't find obvious Tailwind classes"
    echo "   First 10 lines of CSS:"
    cat $CSS_FILES | head -10
fi
echo ""

echo "Step 6: Preview the build..."
echo "Run: cd dist && python3 -m http.server 8000"
echo "Then open: http://localhost:8000"
echo ""

echo "Step 7: Check index.html for CSS link..."
if grep -q "\.css" dist/index.html; then
    echo "✓ index.html references CSS file:"
    grep "\.css" dist/index.html
else
    echo "❌ ERROR: index.html doesn't reference any CSS!"
fi
echo ""

echo "✅ Verification complete!"
echo ""
echo "If everything looks good, build Docker image:"
echo "  docker build -t frontend:test ."
echo ""
echo "Then test it:"
echo "  docker run -p 8085:80 frontend:test"
echo "  Open: http://localhost:8085"
